package response

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(message string, data interface{}) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ResponseError(message string) BaseResponse {
	return BaseResponse{
		Success: false,
		Message: message,
	}
}
