package controller

import(
	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/jinzhu/gorm"
)

type Controller struct {
	DB	*gorm.DB
}

func Init(db *gorm.DB) Controller {
	return Controller{
		DB: db,
	}
}

func (ctrl *Controller) SignInHandler(cxt *gin.Context) {
}

func (ctrl *Controller) SignUpHandler(cxt *gin.Context) {
}
