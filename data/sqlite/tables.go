// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type Task struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name"`
	Done    interface{} `json:"done"`
	Created int64       `json:"created"`
	Type    int64       `json:"type"`
}
