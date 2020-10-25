package controller

import(
	"net/http"
	// import gin library
	"github.com/gin-gonic/gin"

	// Import SQL driver
	"github.com/miraikeitai2020/backend-auth/pkg/hash"
	"github.com/miraikeitai2020/backend-auth/pkg/database"
	"github.com/miraikeitai2020/backend-auth/pkg/server/view"
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
	id := cxt.GetHeader("id")
	if id == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, "id value is empty."),
		)
		return
	}
	pass := cxt.GetHeader("password")
	if pass == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, "password value is empty."),
		)
		return
	}
	user := ctrl.DB.GetUserInfo(id)
	if user.ID == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, id + " user is already exist."),
		)
		return
	}
	if user.Pass != hash.CreateHashString(pass) {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, "password is not true."),
		)
		return
	}
	cxt.Writer.Header().Set("user-id", id)
	cxt.Next()
	cxt.Writer.WriteHeader(http.StatusOK)
}

func (ctrl *Controller) SignUpHandler(cxt *gin.Context) {
	id := cxt.GetHeader("id")
	if id == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, "id value is empty."),
		)
		return
	}
	pass := cxt.GetHeader("password")
	if pass == "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, "password value is empty."),
		)
		return
	}
	user := ctrl.DB.GetUserInfo(id)
	if user.ID != "" {
		cxt.JSON(
			http.StatusBadRequest,
			view.MakeErrorResponse(http.StatusBadRequest, id + " user is already exist."),
		)
		return
	}
	ctrl.DB.InsertUserInfo(
		id, 
		hash.CreateHashString(pass),
	)
	cxt.Writer.Header().Set("user-id", id)
	cxt.Next()
	cxt.Writer.WriteHeader(http.StatusOK)
}
