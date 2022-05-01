package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

// ValidError 验证异常的结构体内容
type ValidError struct {
	Key     string
	Message string
}

// ValidErrors 错误原因有很多的时候，放入结构体切片中
type ValidErrors []*ValidError

// Error 单个错误，返回错误信息
func (v ValidError) Error() string {
	return v.Message
}

// Error 多个错误，返回拼装成字符串的错误信息
func (v ValidErrors) Error() string {
	// 将字符串切片组装成 , 分割的字符串
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	// 声明一个变量-字符串类型的切片
	var errs []string
	// 遍历验证异常的错误
	for _, err := range v {
		// 将每一个错误的字符串追加到这个空切片中
		errs = append(errs, err.Error())
	}
	// 返回这个字符串切片
	return errs
}

func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	// 声明一个变量-验证错误的结构体
	var errs ValidErrors
	// gin 提供的方法，将请求的参数绑定到指定的结构体上
	err := c.ShouldBind(v)
	// 如果绑定的过程中出现了错误
	if err != nil {
		// 获取错误的内容
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
