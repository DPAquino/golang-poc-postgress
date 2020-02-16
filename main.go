package main

import (
	"fmt"

	"golang-poc-postgress/config"
	"golang-poc-postgress/infrastructure/datastore"
	"golang-poc-postgress/infrastructure/router"
	"golang-poc-postgress/registry"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r)

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
