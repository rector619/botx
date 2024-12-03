package external_models

type ValidateOnDBReq struct {
	Table string                 `json:"table"`
	Type  string                 `json:"type"`
	Query map[string]interface{} `json:"query"`
}

type ValidateOnDBReqModel struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}

type ValidateAuthorizationReq struct {
	Type               string `validate:"required" json:"type"`
	AuthorizationToken string `json:"authorization-token"`
	AppKey             string `json:"app-key"`
	PrivateKey         string `json:"private-key"`
	PublicKey          string `json:"public-key"`
}

type ValidateAuthorizationModel struct {
	Status  string                         `json:"status"`
	Code    int                            `json:"code"`
	Message string                         `json:"message"`
	Data    ValidateAuthorizationDataModel `json:"data"`
}
type ValidateAuthorizationDataModel struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
