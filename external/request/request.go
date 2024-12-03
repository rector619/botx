package request

import (
	"fmt"

	"github.com/SineChat/bot-ms/external/microservice/auth"
	"github.com/SineChat/bot-ms/external/microservice/notification"
	"github.com/SineChat/bot-ms/external/mocks"
	"github.com/SineChat/bot-ms/external/thirdparty/ipstack"
	"github.com/SineChat/bot-ms/external/thirdparty/whatsapp"
	"github.com/SineChat/bot-ms/internal/config"
	"github.com/SineChat/bot-ms/utility"
)

type ExternalRequest struct {
	Logger *utility.Logger
	Test   bool
}

var (
	JsonDecodeMethod    string = "json"
	PhpSerializerMethod string = "phpserializer"

	// microservice
	SendWelcomeMail       string = "send_welcome_mail"
	SendResetPasswordMail string = "send_reset_password_mail"
	GetUserReq            string = "get_user"
	ValidateOnAuth        string = "validate_on_auth"
	ValidateAuthorization string = "validate_authorization"
	GetAccessTokenByKey   string = "get_access_token_by_key"

	// third party
	IpstackResolveIp string = "ipstack_resolve_ip"

	// whatsapp
	WhatsAppSendMessage string = "whatsapp_send_message"
)

func (er ExternalRequest) SendExternalRequest(name string, data interface{}) (interface{}, error) {
	var (
		config = config.GetConfig()
	)
	if !er.Test {
		switch name {
		case IpstackResolveIp:
			obj := ipstack.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v", config.IPStack.BaseUrl),
				Method:       "GET",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.IpstackResolveIp()
		case SendWelcomeMail:
			obj := notification.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v1/send/%s", config.Microservices.Notification, SendWelcomeMail),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.SendWelcomeMail()
		case SendResetPasswordMail:
			obj := notification.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v1/send/%s", config.Microservices.Notification, SendResetPasswordMail),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.SendResetPasswordMail()
		case GetUserReq:
			obj := auth.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v1/get_user", config.Microservices.Auth),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.GetUser()
		case ValidateOnAuth:
			obj := auth.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v1/validate_on_db", config.Microservices.Auth),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.ValidateOnAuth()
		case ValidateAuthorization:
			obj := auth.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v1/validate_authorization", config.Microservices.Auth),
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.ValidateAuthorization()
		case GetAccessTokenByKey:
			obj := auth.RequestObj{
				Name:         name,
				Path:         fmt.Sprintf("%v/v2/get_access_token_by_key", config.Microservices.Auth),
				Method:       "GET",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.GetAccessTokenByKey()
		case WhatsAppSendMessage:
			obj := whatsapp.RequestObj{
				Name:         name,
				Path:         "",
				Method:       "POST",
				SuccessCode:  200,
				DecodeMethod: JsonDecodeMethod,
				RequestData:  data,
				Logger:       er.Logger,
			}
			return obj.SendWhatsappMessage()
		default:
			return nil, fmt.Errorf("request not found")
		}

	} else {
		mer := mocks.ExternalRequest{Logger: er.Logger, Test: true}
		return mer.SendExternalRequest(name, data)
	}
}
