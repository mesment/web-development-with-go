package controllers

import (
	"lenslocked.com/views"
)

type Static struct {
	Home    *views.View
	Contact *views.View
	FAQ     *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap-4", "static/home"),
		Contact: views.NewView("bootstrap-4", "static/contact"),
		FAQ:     views.NewView("bootstrap-4", "static/faq"),
	}
}