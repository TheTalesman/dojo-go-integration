package main

import "fmt"

func Main(input string, db DBAdapter, queue QueueAdapter) bool {
	saved, err := db.SaveToDB(&input)
	if err != nil || saved == false {
		return false
	}
	fmt.Println(saved)
	published, err := queue.saveToQueue(&input)
	if err != nil || published == false {
		return false
	}
	fmt.Println(published)
	return true

}

type DBAdapter interface {
	SaveToDB(s *string) (bool, error)
}

type QueueAdapter interface {
	saveToQueue(s *string) (bool, error)
}
