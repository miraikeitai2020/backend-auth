package server

import(
	// import gin library
	"github.com/gin-gonic/gin"

	"github.com/miraikeitai2020/backend-auth/pkg/server/controller"
)

func Router(ctrl controller.Controller) *gin.Engine {
	router := gin.Default()
	router.POST("/query/signin", ctrl.SignInHandler)
	router.POST("/mutation/signup", ctrl.SignUpHandler)
	return router
}