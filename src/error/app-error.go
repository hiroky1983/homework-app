package apperror

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Trace	  string `json:"trace,omitempty"`
}

func ErrorWrapperWithCode(err error, code int) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}
