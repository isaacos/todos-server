package controllers

import (
	"fmt"
	"net/http"

	"todos-backend/models"
)

//NewLists create newList controller
func NewLists(lis *models.ListService) *Lists {
	return &Lists{
		lis: lis,
	}
}

type Lists struct {
	lis *models.ListService
}

//Renders GET /signup form
//func (l *Lists) New(w http.ResponseWriter, r *http.Request) {
//	if err := l.NewView.Render(w, nil); err != nil {
//		panic(err)
//	}
//}

type SignupForm struct {
	Title string `schema:"title" json:"title"`
}

//POST /signup
func (l *Lists) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parsePost(r, w, &form); err != nil {
		panic(err)
	}

	fmt.Println("form ", form)
	list := models.List{
		Title: form.Title,
	}

	if err := l.lis.Create(&list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, form)
}
