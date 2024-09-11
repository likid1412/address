package main

import (
	"fmt"
	"os"

	"github.com/likid1412/address/logger"
	"github.com/likid1412/address/routes"
	"github.com/rs/zerolog/log"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := logger.Init(); err != nil {
		err = fmt.Errorf("logger init failed: %w", err)
		return err
	}

	httpServer := routes.Init()
	if err := httpServer.Run(":8080"); err != nil {
		err = fmt.Errorf("httpServer run failed: %w", err)
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}
