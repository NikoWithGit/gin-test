package main

import (
	"gin-test/api"
	"gin-test/env"
)

func main() {
	/*err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal(err)
	}
	*/
	srv := &api.Server{}

	dsn := env.GetDbDsn()

	srv.InitDb(dsn)
	srv.InitGin()

	srv.RegisterRoutes()

	srv.Start(":8080")
}
