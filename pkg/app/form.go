package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// 对应每个 form 表单的字段验证信息
type ValidatorError struct {
	Key     string
	Message string
}

func (v *ValidatorError) Error() string {
	return v.Message
}

type ValidatorErrors []*ValidatorError

func (v ValidatorErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// 获取所有验证错误信息
func (v ValidatorErrors) Errors() []string {
	var errors []string
	for _, err := range v {
		errors = append(errors, err.Error())
	}

	return errors
}

// 对数据绑定并且验证
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidatorErrors) {
	var errs ValidatorErrors
	// 从 c 中解析字段,并绑定到 v 上面.并且对数据进行验证
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		validErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			errs = append(errs, &ValidatorError{
				Key:     "unknown",
				Message: "unknown",
			})
			return false, errs
		}

		for k, v := range validErrors.Translate(trans) {
			errs = append(errs, &ValidatorError{
				Key:     k,
				Message: v,
			})
		}

		return false, errs
	}

	return true, nil
}
