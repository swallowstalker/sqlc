// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package crud

import (
	"liucai.com/zjbmis/schema"
)

type AuthUser struct {
	ID       int32              `json:"id"`
	Name     string             `json:"name"`
	UserType string             `json:"user_type"`
	Extent   schema.FundSummary `json:"extent"`
	DeptID   *int32             `json:"dept_id"`
}
