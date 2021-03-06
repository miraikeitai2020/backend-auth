package main

import(
	"log"
	"github.com/miraikeitai2020/backend-auth/pkg/config"
	"github.com/miraikeitai2020/backend-auth/pkg/database"
	"github.com/miraikeitai2020/backend-auth/pkg/server"
	"github.com/miraikeitai2020/backend-auth/pkg/server/controller"
)

func main() {
	addr := config.GetRouterAddr()

	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	ctrl := controller.Init(db)
	if err := server.Router(ctrl).Run(addr); err != nil {
		panic(err)
	}
}