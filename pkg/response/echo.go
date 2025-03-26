package response

import "github.com/labstack/echo/v4"

func JSON(c echo.Context, status APIStatus, data interface{}) error {
	var response interface{}

	if status.Code >= 200 && status.Code < 300 {
		response = ResponseSuccess(status.Message, data)
	} else if status.Code >= 400 && status.Code < 500 {
		response = ResponseError(status.Message)
	} else {
		response = ResponseError(status.Message)
	}

	return c.JSON(status.Code, response)
}
