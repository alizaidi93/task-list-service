package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	task "task-list-service/ayzee/database/api"
	"task-list-service/ayzee/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var requestCounter int = 0

type Handlers struct {
	service *service.Service
	log     *logrus.Logger
}

func ProvideHandlers(service *service.Service, log *logrus.Logger) *Handlers {
	return &Handlers{service: service, log: log}
}

func (h Handlers) GetTask(w http.ResponseWriter, r *http.Request) {
	// requestCounter := logReqest(r, h.log)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	uuid := params["uuid"]
	task := h.service.GetTask(uuid)
	io.WriteString(w, string(task.ConvertToJson()))
	// logResponse(w, requestCounter, h.log)
}

func (h Handlers) GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	user := params["user"]
	tasks := h.service.GetTasksByUser(user)
	io.WriteString(w, task.ConvertTaskSliceToJson(tasks))
}

func (h Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBytes, _ := ioutil.ReadAll(r.Body)
	var task task.Task
	json.Unmarshal(requestBytes, &task)
	uuid := h.service.CreateTask(task)
	io.WriteString(w, uuid)
}

func (h Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	requestBytes, _ := ioutil.ReadAll(r.Body)
	var task task.Task
	json.Unmarshal(requestBytes, &task)
	h.service.UpdateTask(task)
	w.WriteHeader(http.StatusNoContent)
}

//Helper Method for local debugging
func logReqest(httpRequest *http.Request, log *logrus.Logger) int {
	requestCounter++
	headers := httpRequest.Header
	// Orders the Header names alphabetically
	alphabetisedKeys := make([]string, 0, len(headers))
	for k := range headers {
		alphabetisedKeys = append(alphabetisedKeys, k)
	}
	sort.Strings(alphabetisedKeys)

	//Logs the Headers of an incoming request
	var alphabetisedHeaders bytes.Buffer
	for k := range alphabetisedKeys {
		alphabetisedHeaders.WriteString(fmt.Sprintf("%s:%s", alphabetisedKeys[k], headers[alphabetisedKeys[k]]))
	}
	log.Info(alphabetisedHeaders.String())

	return requestCounter
}

//Helper Method for local debugging
func logResponse(w http.ResponseWriter, currentCounter int, log *logrus.Logger) {
	responseHeaders := w.Header()
	log.Info("Response Sent From Server")
	for i, v := range responseHeaders {
		log.Info(fmt.Sprintf("Index : %s Value : %s", i, v))
	}

}
