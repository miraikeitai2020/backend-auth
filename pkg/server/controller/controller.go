package controller

import(
	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/miraikeitai2020/backend-auth/pkg/database"
)

type Controller struct {
	DB	*database.DB
}

func Init(db *database.DB) Controller {
	return Controller{
		DB: db,
	}
}

func (ctrl *Controller) SignInHandler(cxt *gin.Context) {
}

func (ctrl *Controller) SignUpHandler(cxt *gin.Context) {
}
