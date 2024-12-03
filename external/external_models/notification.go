package external_models

type SendResetPasswordMail struct {
	Email string `json:"email"`
	Token int    `json:"token"`
}

type SendWelcomeMail struct {
	Email string `json:"email"`
}
