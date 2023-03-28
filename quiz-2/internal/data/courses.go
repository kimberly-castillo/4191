// Filename: cmd/api/courses.go
package data

import (
	"time"
)

// Course represents one row of data in our schools table
type Course struct {
	ID       int64     `json:"id"`
	Code     string    `json:"course_code"`
	Title    string    `json:"course_title"`
	Credits  string    `json:"credits"`
	CreateAt time.Time `json:"-"`
	Version  int32     `json:"version"`
}
