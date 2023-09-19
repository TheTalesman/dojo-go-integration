package adapter

import (
	"errors"
)

type DBAdapter struct {
}

// saveToDB implements main.DBAdapter.
func (DBAdapter) SaveToDB(s *string) (bool, error) {
	if *s == "correto" {
		return true, nil
	}
	return false, errors.New("could not save to db")
}

// func (dbMock DBMock) saveToDB(s *string) (bool, error) {
// 	return true, nil
// }
