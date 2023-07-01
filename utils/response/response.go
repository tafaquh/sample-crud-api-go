package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseMap(data interface{}, message string) *Response {
	return &Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
