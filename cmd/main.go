package main

import (
	"fmt"
	"os"

	"github.com/Jhon-2801/cli-todo/core/task"
	"github.com/Jhon-2801/cli-todo/db"
)

func main() {

	db, err := db.BDConnection()

	if err != nil {
		fmt.Println(err)
	}

	ser := task.NewSer(db)
	if len(os.Args) < 2 {
		prinUsage()
		return
	}

	switch os.Args[1] {
	case "list":
		ser.GetAll()
	case "add":
		ser.Create()
	case "delete":
		ser.Delete()
	case "complete":
		ser.Complete()
	default:
		prinUsage()
	}

}

func prinUsage() {
	fmt.Println("Uso: cli [list|add|complete|delete]")
}
