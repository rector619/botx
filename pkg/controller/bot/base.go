package bot

import (
	"github.com/SineChat/bot-ms/external/request"
	"github.com/SineChat/bot-ms/pkg/repository/storage/mongodb"
	"github.com/SineChat/bot-ms/utility"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Db        *mongodb.Database
	Validator *validator.Validate
	Logger    *utility.Logger
	ExtReq    request.ExternalRequest
}
