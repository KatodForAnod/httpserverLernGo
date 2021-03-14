package dto

import "fmt"

// UserStruct - dto user
type UserStruct struct {
	UserName string
	UserID   int
}

// UserStructToString - get all information from dto user in string
func (s UserStruct) UserStructToString() string {
	return fmt.Sprintf("%s - ID: %d\n", s.UserName, s.UserID)
}
