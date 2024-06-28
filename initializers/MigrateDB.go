package initializers

import "groq-api/db"

func MigrateDB() {
	err := DB.AutoMigrate(&db.ClientUsers{}, &db.History{})
	if err != nil {
		return
	}
}
