package helper

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
)

func JSONResponse(c echo.Context, statusCode int, msg interface{}) error {

	return c.JSON(statusCode, map[string]interface{}{
		"dto": msg,
	})
}
func RespError(msg interface{}) interface{} {

	return  map[string]interface{}{
		"err": msg,
	}
}

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}

	return string(b)
}

func Validator(strc interface{}) []string {
	validate := validator.New()
	var totalErrors []string
	err := validate.Struct(strc)
	if err != nil {
		
		for i, err := range err.(validator.ValidationErrors) {
			theError := fmt.Sprintf("Validation error number %d for %s: %s", i, err.Field(), err.Tag())
			totalErrors = append(totalErrors, theError)
		}
		return totalErrors
	}
	return totalErrors
}

