package cmd

import (
	"fmt"
	"strconv"

	"github.com/forzeyy/todo-cli-app/database"
	"github.com/forzeyy/todo-cli-app/model"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [taskid]",
	Short: "mark a task as done",
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
		}
		task.Status = true
		database.Database.Db.Save(&task)
		fmt.Printf("task with ID %d marked as done\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
