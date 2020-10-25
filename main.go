package main

import(
	"github.com/miraikeitai2020/backend-auth/pkg/config"
	"github.com/miraikeitai2020/backend-auth/pkg/server"
	"github.com/miraikeitai2020/backend-auth/pkg/server/controller"
)

func main() {
	addr := config.GetRouterAddr()

	// TODO: database connection function exec here.

	ctrl := controller.Init(nil)
	if err := server.Router(ctrl).Run(addr); err != nil {
		panic(err)
	}
}