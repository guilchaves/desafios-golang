package response

type Response struct {
	Status int    `json:"status,omitempty"`
	Error  string `json:"message,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func ErrorResponse(status int, message string) *Response {
	return &Response{
		Status: status,
		Error:  message,
	}
}
