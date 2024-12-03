package bot

import (
	"fmt"

	"github.com/SineChat/bot-ms/internal/models"
)

func ValidateCreateActionTypes(req models.AddActionReq) error {

	switch req.Type {
	case models.TextActionType:
		return ValidateCreateTextActionTypes(req)
	default:
		return fmt.Errorf("not  implemented")
	}
}

func ValidateCreateTextActionTypes(req models.AddActionReq) error {

	if req.Text == nil {
		return fmt.Errorf("missing text object")
	}

	if req.Text.Body == "" {
		return fmt.Errorf("text.body is required")
	}

	return nil
}
