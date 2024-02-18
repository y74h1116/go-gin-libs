package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// custom validation
const CHARACTERS = `^[0-9a-zA-Z\!\?\-]+$`
func isPasswordCharacters(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if ok == false {
		return false
	}
	result := regexp.MustCompile(CHARACTERS).MatchString(password)
	return result
}

func main() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("custom_password_characters", isPasswordCharacters)
	}

	router := gin.Default()

	router.Static("assets", "resource/assets")
	router.LoadHTMLGlob("resource/view/*/*")

	login := NewLogin()
	router.GET("/", login.index)
	router.POST("/", login.login)

	home := NewHome()
	router.GET("/home", home.index)

	router.Run(":8080")
}
