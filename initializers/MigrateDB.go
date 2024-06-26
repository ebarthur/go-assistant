package initializers

import "groq-api/db"

func MigrateDB() {
	err := DB.AutoMigrate(&db.ClientUsers{})
	if err != nil {
		return
	}
}
