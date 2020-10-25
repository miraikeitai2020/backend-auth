package view

import(
	"github.com/miraikeitai2020/backend-auth/pkg/server/model"
)

const(
	STATUS_CODE_400 = "Bad Request"
	STATUS_CODE_403 = "Forbidden"
	STATUS_CODE_500 = "Internal Server Error"
)

func MakeErrorResponse(code int, desc string) model.Error {
	var msg string
	switch code {
	case 400:
		msg = STATUS_CODE_400
	case 403:
		msg = STATUS_CODE_403
	case 500:
		msg = STATUS_CODE_500
	}
	return model.Error{
		Code: code,
		Message: msg,
		Description	: desc,
	}
}
