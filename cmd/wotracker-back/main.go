package main

import (
	log "github.com/sirupsen/logrus"
	"wotracker-back/internal/app/starter"
)

// @title           Bootstrap Iris
// @version         1.0
// @description     This is a sample iris server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4242
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationUrl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("starting API... ")
	starter.RunApi()
	log.Info("ending API... ")
}
