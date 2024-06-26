package main

import (
	"groq-api/cmd/api"
	"groq-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	router := api.SetupRouter()

	err := router.Run()
	if err != nil {
		panic(err)
		return
	}
}
