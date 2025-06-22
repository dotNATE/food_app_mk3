package utils

type HTTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func CreateSuccessfulHTTPResponse(message string, data any) HTTPResponse {
	return HTTPResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func CreateErrorHTTPResponse(message string, err error) HTTPResponse {
	if err != nil {
		message = message + err.Error()
	}

	return HTTPResponse{
		Success: false,
		Error:   message,
	}
}
