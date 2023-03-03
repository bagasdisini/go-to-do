package util

import (
	"github.com/go-playground/validator/v10"
	"go-to-do/dto"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []*dto.ErrorResponse {
	var errors []*dto.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ErrorResponse
			var errorField string
			if err.Field() == "Title" {
				errorField = "title"
			} else if err.Field() == "ActivityGroupID" {
				errorField = "activity_group_id"
			}

			element.FailedField = errorField
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
