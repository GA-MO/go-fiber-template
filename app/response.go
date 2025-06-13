package app

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse(message string, data any) Response {
	return Response{
		Status:  StatusSuccess,
		Message: message,
		Data:    data,
	}
}

func NewResponseError(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}
