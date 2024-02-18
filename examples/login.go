package main

import (
	"net/http"

	"github.com/y74h1116/go-gin-libs/page"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (login *Login) index(ctx *gin.Context) {
	viewParam := login.pageLogin.CreateViewParam()
	ctx.HTML(http.StatusOK, "login/index", viewParam)
}

type LoginForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required,custom_password_characters"`
}

var loginValidationMessages = map[string]map[string]string{
	"Email": {
		"required": "メールアドレスを入力してください。",
	},
	"Password": {
		"required": "パスワードを入力してください。",
		"custom_password_characters": "パスワードに利用できない文字が含まれています。",
	},
}

const EMAIL = "aaa123@aaa.com"
const PASSWORD = "abcd1234"

func (login *Login) login(ctx *gin.Context) {

	var loginForm LoginForm

	if err := ctx.ShouldBind(&loginForm); err != nil {
		viewParam := login.pageLogin.CreateViewParam()
		login.pageLogin.SetFormErrorsByValidator(viewParam, err.(validator.ValidationErrors))
		login.pageLogin.SetOld(viewParam, loginForm)
		ctx.HTML(http.StatusBadRequest, "login/index", viewParam)
		return
	}

	if (loginForm.Email != EMAIL || loginForm.Password != PASSWORD) {
		viewParam := login.pageLogin.CreateViewParam()
		login.pageLogin.SetErrorMessage(viewParam, "メールアドレスまたはパスワードを間違えています。")
		login.pageLogin.SetOld(viewParam, loginForm)
		ctx.HTML(http.StatusBadRequest, "login/index", viewParam)
		return		
	}

	ctx.Redirect(http.StatusFound, "/home")
}

type Login struct {
	pageLogin *page.Page
}

func NewLogin() *Login {
	pageLogin := page.NewPage("ログイン", loginValidationMessages)
	return &Login{pageLogin}
}
