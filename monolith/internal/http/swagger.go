package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mrcampbell/gacha-game/monolith/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gacha Game: Monolith
// @version 1.0
// @description Proof of Concept, inmemory Gacha Game

// // @contact.name API Support
// // @contact.url http://www.swagger.io/support
// // @contact.email support@swagger.io

// // @license.name Apache 2.0
// // @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// // @securityDefinitions.basic BasicAuth

// // @securityDefinitions.apikey ApiKeyAuth
// // @in header
// // @name Authorization

// // @securitydefinitions.oauth2.application OAuth2Application
// // @tokenUrl https://example.com/oauth/token
// // @scope.write Grants write access
// // @scope.admin Grants read and write access to administrative information
//
// // @securitydefinitions.oauth2.implicit OAuth2Implicit
// // @authorizationurl https://example.com/oauth/authorize
// // @scope.write Grants write access
// // @scope.admin Grants read and write access to administrative information
//
// // @securitydefinitions.oauth2.password OAuth2Password
// // @tokenUrl https://example.com/oauth/token
// // @scope.read Grants read access
// // @scope.write Grants write access
// // @scope.admin Grants read and write access to administrative information
//
// // @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// // @tokenUrl https://example.com/oauth/token
// // @authorizationurl https://example.com/oauth/authorize
// // @scope.admin Grants read and write access to administrative information
//
// // @x-extension-openapi {"example": "value on a json format"}

func registerSwaggerHandlers(r *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
