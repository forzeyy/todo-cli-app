package cmd

import (
	"fmt"
	"os"

	"github.com/forzeyy/todo-cli-app/database"
	"github.com/forzeyy/todo-cli-app/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		var undoneTasks uint
		var tasks []model.Task
		database.Database.Db.Order("priority desc").Find(&tasks)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "status", "task name", "description", "priority", "created"})
		for _, task := range tasks {
			status := " "
			if task.Status {
				status = "✓"
			} else {
				undoneTasks += 1
			}
			t.AppendRow(table.Row{task.ID, status, task.Name, task.Description, task.Priority, task.CreatedAt.Format("2006-01-02 15:04:05")})
		}
		t.SetTitle(fmt.Sprintf("you have %d undone tasks!", undoneTasks))
		t.Render()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
