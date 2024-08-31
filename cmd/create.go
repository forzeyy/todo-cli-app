package cmd

import (
	"fmt"
	"time"

	"github.com/forzeyy/todo-cli-app/database"
	"github.com/forzeyy/todo-cli-app/model"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [name] [description] [priority]",
	Short: "add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		priority := 1
		fmt.Sscanf(args[2], "%d", &priority)

		task := model.Task{
			Name:        args[0],
			Description: args[1],
			Priority:    priority,
			CreatedAt:   time.Now(),
			Status:      false,
		}
		database.Database.Db.Create(&task)
		fmt.Printf("task \"%s\" added with priority %d!\n", task.Name, task.Priority)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
