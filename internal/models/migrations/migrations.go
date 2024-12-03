package migrations

import "github.com/SineChat/bot-ms/internal/models"

// _ = db.AutoMigrate(MigrationModels()...)
func AuthMigrationModels() []interface{} {
	return []interface{}{
		&models.Action{},
		&models.Bot{},
		&models.Connection{},
		&models.Integration{},
		&models.WebhookLog{},
	}
}
