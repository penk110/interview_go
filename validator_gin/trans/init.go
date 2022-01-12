package trans

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Trans is a global translator
var Trans ut.Translator

// Init is init trans and return an error when init failed
func Init(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// zhT enT
		zhT := zh.New()
		enT := en.New()

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// v.RegisterStructValidation(ParamStructLevelValidation, UserInfo{})

		uni := ut.New(enT, zhT, enT)
		// locale deside to http header params - 'Accept-Language'
		var ok bool
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("[trans.Init] get translator failed, locale: %s", locale)
		}
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return nil
}

// RemoveTopStruct remove struct tag
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// ParamStructLevelValidation 自定义结构体校验函数
func ParamStructLevelValidation(sl validator.StructLevel) {
	// su := sl.Current().Interface().(UserInfo)

	// if su.Password != su.RePassword {
	// 	// 输出错误提示信息
	// 	sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	// }
}
