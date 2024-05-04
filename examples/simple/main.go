package main

import (
	"fmt"
	"github.com/fumeapp/taskin"
	"time"
)

func main() {

	tasks := taskin.Taskin(taskin.Tasks{
		{
			Title: "Task 1",
			// sleep for 3 seconds then return nil
			Task: func(t *taskin.Task) error {
				for i := 0; i < 3; i++ {
					t.Title = fmt.Sprintf("Task 1: [%d/3] second has passed", i+1)
					time.Sleep(1 * time.Second)
				}
				return nil
			},
		},
		{
			Title: "Task 2",
			// sleep for 3 seconds then return nil
			Task: func(t *taskin.Task) error {
				for i := 0; i < 3; i++ {
					t.Title = fmt.Sprintf("Task 1: [%d/3] second has passed", i+1)
					time.Sleep(1 * time.Second)
				}
				return nil
			},
		},
	})
	err := tasks.Run()

	if err != nil {
		panic(err)
	}
}
