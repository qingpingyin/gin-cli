package serror

const (
	// server error
	Success         = 0
	NotFound        = 404
	InternalServer  = 500
	BadParamRequest = 400

	//service error
	UserNotExist = 1001
)

var (
	ErrSuccess         = NewResponseError(Success, "ok")
	ErrNotFound        = NewResponseError(NotFound, "not found")
	ErrInternalServer  = NewResponseError(InternalServer, "internal server error")
	ErrBadParamRequest = NewResponseError(BadParamRequest, "bad request")
)

var (
	ErrUserNotExist = NewResponseError(UserNotExist, "the user does not exist")
)
