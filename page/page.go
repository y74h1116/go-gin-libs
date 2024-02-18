package page

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (page *Page) CreateViewParam() gin.H {
	viewParam := gin.H{}
	viewParam["pageTitle"] = page.PageTitle
	return viewParam
}

func (page *Page) SetErrorMessage(viewParam gin.H, errorMessage string) {
	viewParam["errorMessage"] = errorMessage
}

func (page *Page) SetFormErrorsByValidator(viewParam gin.H, ve validator.ValidationErrors) {
	formErrors := map[string]string{}
	for _, e := range ve {
		if str, ok := page.ValidationMessages[e.Field()][e.Tag()]; ok {
			formErrors[e.Field()] = str
		}
	}
	viewParam["formErrors"] = formErrors
}

func (page *Page) SetFormErrors(viewParam gin.H, formErrors map[string]string) {
	viewParam["formErrors"] = formErrors
}

func (page *Page) SetOld(viewParam gin.H, old any) {
	viewParam["old"] = old
}

func (page *Page) SetData(viewParam gin.H, datas any) {
	viewParam["datas"] = datas
}

type Page struct {
	PageTitle          string
	ValidationMessages map[string]map[string]string
}

func NewPage(pageTitle string, validationMessages map[string]map[string]string) *Page {
	return &Page{pageTitle, validationMessages}
}
