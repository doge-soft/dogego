package utils

import (
	"{{cookiecutter.app_name}}/api/serializer"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

// 参数错误返回
func ParamaterErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := fmt.Sprintf("Field.%s", e.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)
			return serializer.ParamaterErrorResponse(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamaterErrorResponse("JSON类型不匹配", err)
	}

	return serializer.ParamaterErrorResponse("参数错误", err)
}
