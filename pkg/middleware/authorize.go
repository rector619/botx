package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SineChat/bot-ms/external/external_models"
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/internal/config"
	"github.com/SineChat/bot-ms/internal/models"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/utility"
	"github.com/gin-gonic/gin"
)

const (
	AuthType       AuthorizationType = "auth"
	ApiPublicType  AuthorizationType = "api_public"
	ApiPrivateType AuthorizationType = "api_private"
	AppType        AuthorizationType = "app"
)

type (
	AuthorizationType  string
	AuthorizationTypes []AuthorizationType
)

func Authorize(db *mongodb.Database, extReq request.ExternalRequest, authTypes ...AuthorizationType) gin.HandlerFunc {

	return func(c *gin.Context) {
		if len(authTypes) > 0 {

			msg := ""
			for _, v := range authTypes {
				ms, status := v.ValidateAuthorizationRequest(c, extReq, db)
				if status {
					return
				}
				msg = ms
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, utility.UnauthorisedResponse(http.StatusUnauthorized, fmt.Sprint(http.StatusUnauthorized), "Unauthorized", msg))
		}
	}
}

func (at AuthorizationType) ValidateAuthorizationRequest(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database) (string, bool) {
	if at == ApiPublicType {
		return at.ValidateApiPublicType(c, extReq, db)
	} else if at == ApiPrivateType {
		return at.ValidateApiPrivateType(c, extReq, db)
	} else if at == AppType {
		return at.ValidateAppType(c, extReq)
	} else if at == AuthType {
		return at.ValidateAuthType(c, extReq, db)
	}

	return "authorized", true
}

func (at AuthorizationType) ValidateAuthType(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database) (string, bool) {

	var invalidToken = "Your request was made with invalid credentials."
	authorizationToken := GetHeader(c, "Authorization")
	if authorizationToken == "" {
		return "token not provided", false
	}

	bearerTokenArr := strings.Split(authorizationToken, " ")
	if len(bearerTokenArr) != 2 {
		return invalidToken, false
	}

	bearerToken := bearerTokenArr[1]

	if bearerToken == "" {
		return invalidToken, false
	}

	reqInf, err := extReq.SendExternalRequest(request.ValidateAuthorization, external_models.ValidateAuthorizationReq{
		Type:               string(AuthType),
		AuthorizationToken: bearerToken,
	})
	if err != nil {
		return err.Error(), false
	}

	dataResponse := reqInf.(external_models.ValidateAuthorizationDataModel)
	if !dataResponse.Status {
		return dataResponse.Message, false
	}

	user := external_models.User{}

	uByte, err := json.Marshal(dataResponse.Data)
	if err != nil {
		return err.Error(), false
	}

	err = json.Unmarshal(uByte, &user)
	if err != nil {
		return err.Error(), false
	}

	models.MyIdentity = &user

	return "authorized", true
}

func (at AuthorizationType) ValidateAppType(c *gin.Context, extReq request.ExternalRequest) (string, bool) {
	config := config.GetConfig().App
	appKey := GetHeader(c, "app-key")
	if appKey == "" {
		return "missing app key", false
	}

	if appKey != config.Key {
		return "invalid app key", false
	}

	return "authorized", true
}

func (at AuthorizationType) ValidateApiPublicType(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database) (string, bool) {
	publicKey := GetHeader(c, "public-key")

	if publicKey == "" {
		return "missing api key", false
	}

	reqInf, err := extReq.SendExternalRequest(request.ValidateAuthorization, external_models.ValidateAuthorizationReq{
		Type:      string(ApiPublicType),
		PublicKey: publicKey,
	})
	if err != nil {
		return err.Error(), false
	}

	dataResponse := reqInf.(external_models.ValidateAuthorizationDataModel)
	if !dataResponse.Status {
		return dataResponse.Message, false
	}

	token := external_models.AccessToken{}

	tByte, err := json.Marshal(dataResponse.Data)
	if err != nil {
		return err.Error(), false
	}

	err = json.Unmarshal(tByte, &token)
	if err != nil {
		return err.Error(), false
	}

	models.MyAccessToken = &token

	return "authorized", true
}
func (at AuthorizationType) ValidateApiPrivateType(c *gin.Context, extReq request.ExternalRequest, db *mongodb.Database) (string, bool) {
	privateKey := GetHeader(c, "private-key")

	if privateKey == "" {
		return "missing api key", false
	}

	reqInf, err := extReq.SendExternalRequest(request.ValidateAuthorization, external_models.ValidateAuthorizationReq{
		Type:       string(ApiPrivateType),
		PrivateKey: privateKey,
	})
	if err != nil {
		return err.Error(), false
	}

	dataResponse := reqInf.(external_models.ValidateAuthorizationDataModel)
	if !dataResponse.Status {
		return dataResponse.Message, false
	}

	token := external_models.AccessToken{}

	tByte, err := json.Marshal(dataResponse.Data)
	if err != nil {
		return err.Error(), false
	}

	err = json.Unmarshal(tByte, &token)
	if err != nil {
		return err.Error(), false
	}

	models.MyAccessToken = &token

	return "authorized", true
}

func GetHeader(c *gin.Context, key string) string {
	header := ""
	if c.GetHeader(key) != "" {
		header = c.GetHeader(key)
	} else if c.GetHeader(strings.ToLower(key)) != "" {
		header = c.GetHeader(strings.ToLower(key))
	} else if c.GetHeader(strings.ToUpper(key)) != "" {
		header = c.GetHeader(strings.ToUpper(key))
	} else if c.GetHeader(strings.Title(key)) != "" {
		header = c.GetHeader(strings.Title(key))
	}
	return header
}
