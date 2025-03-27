package server

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *Error      `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
		Error:   nil,
	}
}

func ErrorResponse(code string, message string) Response {
	return Response{
		Success: false,
		Data:    nil,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}

var Success = SuccessResponse(nil)
