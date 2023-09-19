package main

import (
	"dojo/adapter"
	"testing"
)

// type DBMock struct {
// }

// func (dbMock DBMock) saveToDB(s *string) (bool, error) {
// 	return true, nil
// }

type QueueMock struct {
}

func (queueMock QueueMock) saveToQueue(s *string) (bool, error) {
	return true, nil
}

func TestMain(t *testing.T) {
	input := "correto"
	dbAdapter := adapter.DBAdapter{}
	queueMock := QueueMock{}
	funcionou := Main(input, dbAdapter, queueMock)
	if !funcionou {
		t.Error("Não funcionou!")
	}

}
