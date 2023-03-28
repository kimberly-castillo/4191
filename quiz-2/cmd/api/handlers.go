// Filename: cmd/api/handlers.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kimberly-castillo/4191/internal/data"
)

func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	//creating a struct to hold a school that will be provided to us via the request
	//api should be able to hold post requests
	//post (write) means you want to affect the server/db
	//GET doesnt change the db
	//user will send it as a json object, so we take it, store it in go program in instance of struct in order to save to db

	var input struct {
		ID       int64     `json:"id"`
		Code     string    `json:"course_code"`
		Title    string    `json:"course_title"`
		Credits  string    `json:"credits"`
		CreateAt time.Time `json:"-"`
		Version  int32     `json:"version"`
	}
	//decode the JSON request
	err := app.readJSON(w, r, &input)
	if err != nil {
		//decode returns error so check for error
		app.badRequestResponse(w, r, err)
		return
	}
	//print the request
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//fmt.Fprintf(w, "show details of school %d\n", id)
	course := data.Course{
		ID:       id,
		CreateAt: time.Now(),
		Code:     "CMPS4191",
		Title:    "Advanced Web",
		Credits:  "3",
		Version:  1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
