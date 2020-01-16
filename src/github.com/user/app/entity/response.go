package entity

// Response -
type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Status -
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SetResponse -
func SetResponse(code int, message string, data interface{}) *Response {
	res := &Response{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
	return res
}
