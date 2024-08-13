package main

import (
	"rest_api_sklad_project/cmd/apiserver"
	_ "rest_api_sklad_project/docs"
)

// @title			Sklad API Documentation
// @version		1.0
// @description	Sklad API Documentation
// @termsOfService	http://swagger.io/terms/
func main() {
	apiserver.HandleRequests()
}
