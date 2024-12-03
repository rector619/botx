package external_models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserRequestModel struct {
	ID           string `json:"id"`
	EmailAddress string `json:"email_address"`
}

type GetUserResponse struct {
	Status  string              `json:"status"`
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetUserResponseData `json:"data"`
}
type GetUserResponseData struct {
	User User `json:"user"`
}
