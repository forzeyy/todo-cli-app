package main

import (
	"github.com/forzeyy/todo-cli-app/cmd"
	"github.com/forzeyy/todo-cli-app/database"
)

func main() {
	database.ConnectDb()
	cmd.Execute()
}
