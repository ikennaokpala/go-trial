package controllers

import (
	"net/http"

	"github.com/unrolled/render"

	"bitbucket.org/globalfoodbook/api/dao"
	"bitbucket.org/globalfoodbook/api/models"
)

// Posts controller
type Posts struct{}

// Index action of the Posts controller
func (p Posts) Index(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})

	var posts []models.Post

	dao := dao.Classic()
	dao.Find(&posts)

	r.JSON(res, 200, &posts)
}
