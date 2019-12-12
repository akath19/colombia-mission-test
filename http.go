package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func (s *server) showForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.GetTemplate(w, r)
	}

	if r.Method == http.MethodPost {
		s.ProcessResponses(w, r)
	}
}

func (s *server) GetTemplate(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/css") {
		http.ServeFile(w, r, "html"+r.URL.Path)
	}

	t, err := template.ParseFiles("html/form.html")

	if err != nil {
		log.Fatalf("Cannot parse html template, details: %v", err)
	}

	response := struct {
		Success     bool
		ErrorsFound bool
		Errors      []string
	}{
		Success:     false,
		ErrorsFound: false,
		Errors:      nil,
	}

	t.Execute(w, response)
}

func (s *server) ProcessResponses(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatalf("Cannot parse form, details %v", err)
	}

	name := strings.TrimSpace(r.FormValue("name"))
	color := strings.TrimSpace(r.FormValue("color"))
	cats := r.FormValue("cats")
	dogs := r.FormValue("dogs")

	var errors []string

	if name == "" {
		errors = append(errors, "You must enter your name")
	}

	if color == "" {
		errors = append(errors, "You must enter your favorite color")
	}

	if cats == "" && dogs == "" {
		errors = append(errors, "You must select either cats or dogs")
	}

	var res bool
	var createErr error

	if cats == "on" {
		res, createErr = s.SaveData(name, color, true)
	} else {
		res, createErr = s.SaveData(name, color, false)
	}

	if !res {
		errors = append(errors, "Couldn't save your data to DB, details: "+createErr.Error())
	}

	t, err := template.ParseFiles("html/form.html")

	if err != nil {
		log.Fatalf("Cannot parse html template, details: %v", err)
	}

	var response struct {
		Success     bool
		ErrorsFound bool
		Errors      []string
	}

	if len(errors) > 0 {
		response.Success = false
		response.ErrorsFound = true
		response.Errors = errors
	} else {
		response.Success = true
		response.ErrorsFound = false
		response.Errors = nil
	}

	t.Execute(w, response)
}
