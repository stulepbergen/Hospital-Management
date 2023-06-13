package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Print("main")
	var filename = "././hospitalWeb/main/main.html"
	t, err := template.ParseFiles(filename)
	Error(err)

	err = t.Execute(w, filename)
	Error(err)
}
