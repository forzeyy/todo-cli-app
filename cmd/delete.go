package cmd

import (
	"fmt"
	"strconv"

	"github.com/forzeyy/todo-cli-app/database"
	"github.com/forzeyy/todo-cli-app/model"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("error: invalid ID format")
			return
		}
		var task model.Task
		database.Database.Db.First(&task, id)
		if task.ID == 0 {
			fmt.Printf("error: task with ID %d not found", id)
			return
		}
		database.Database.Db.Delete(&task)
		fmt.Println("task deleted successfully")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
