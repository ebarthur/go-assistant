package main

import (
	"groq-api/cmd/v1/api"
	_ "groq-api/docs"
	"groq-api/initializers"
)

//	@title			            Go-assistant API
//	@version		            1.0
//	@description	            API Documentation for Go-assistant
//	@termsOfService	            http://swagger.io/terms/Go-assistant
//	@contact.name	            Ebenezer Arthur
//	@contact.url	            https://ebarthur.vercel.app
//	@contact.email	            arthurebenezer@aol.com
//	@license.name	            Apache 2.0
//	@license.url	            http://www.apache.org/licenses/LICENSE-2.0.html
//	@host		                localhost:10000
//	@BasePath	                /api/v1
//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

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
	}
}
