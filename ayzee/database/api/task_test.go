package task_test

import (
	"fmt"
	"strings"
	task "task-list-service/ayzee/database/api"
	"testing"
)

func TestConvertToJson(t *testing.T) {
	task := task.Task{1, "uuid", "myTask", true, "AyZee", "AyZee", "2022-01-19 21:56:49.498923"}
	actualJson := task.ConvertToJson()
	expectedJson := "{\"Id\":1,\"Uuid\":\"uuid\",\"Task\":\"myTask\",\"Completed\":true,\"CreatedBy\":\"AyZee\",\"ModifiedBy\":\"AyZee\",\"ModifiedOn\":\"2022-01-19 21:56:49.498923\"}"
	if strings.Compare(actualJson, expectedJson) != 0 {
		t.Fatal("Json incorrectly generated.")
	}
}

func TestConvertTaskSliceToJson(t *testing.T) {
	task1 := task.Task{1, "uuid", "myTask", true, "AyZee", "AyZee", "2022-01-19 21:56:49.498923"}
	task2 := task.Task{1, "uuid", "myTask", true, "AyZee", "AyZee", "2022-01-19 21:56:49.498923"}
	tasks := []task.Task{task1, task2}
	actualJson := task.ConvertTaskSliceToJson(tasks)
	fmt.Println(actualJson)
	expectedJson := "[{\"Id\":1,\"Uuid\":\"uuid\",\"Task\":\"myTask\",\"Completed\":true,\"CreatedBy\":\"AyZee\",\"ModifiedBy\":\"AyZee\",\"ModifiedOn\":\"2022-01-19 21:56:49.498923\"},{\"Id\":1,\"Uuid\":\"uuid\",\"Task\":\"myTask\",\"Completed\":true,\"CreatedBy\":\"AyZee\",\"ModifiedBy\":\"AyZee\",\"ModifiedOn\":\"2022-01-19 21:56:49.498923\"}]"
	if strings.Compare(actualJson, expectedJson) != 0 {
		t.Fatal("Json incorrectly generated.")
	}
}
