package task

import (
	"encoding/json"
)

type Task struct {
	Id         int
	Uuid       string
	Task       string
	Completed  bool
	CreatedBy  string
	ModifiedBy string
	ModifiedOn string
}

func (t Task) ConvertToJson() string {
	barr, _ := json.Marshal(t)
	jsonString := string(barr)
	return jsonString
}

func ConvertTaskSliceToJson(tasks []Task) string {
	barr, _ := json.Marshal(tasks)
	jsonString := string(barr)
	return jsonString
}
