package apperror

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorWrapperWithCode(err error, code int) errorResponse {
	return errorResponse{
		Code:    code,
		Message: err.Error(),
	}
}
