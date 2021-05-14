package handler

import (
	"encoding/json"
	"net/http"
	"rest_demo/pkg/data"
	"rest_demo/pkg/dto"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTodo(writer http.ResponseWriter, req *http.Request) {
	responeWithJson(writer, http.StatusOK, data.Todos)
}

func GetTodoById(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responeWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
	}
	for _, todo := range data.Todos {
		if todo.ID == id {
			responeWithJson(writer, http.StatusOK, todo)
			return
		}
	}
	responeWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
func CreateTodo(writer http.ResponseWriter, req *http.Request) {
	var newTodo dto.Todo
	if err := json.NewDecoder(req.Body).Decode(&newTodo); err != nil {
		responeWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	newTodo.ID = generateId(data.Todos)
	data.Todos = append(data.Todos, newTodo)

	responeWithJson(writer, http.StatusOK, newTodo)

}
func UpdateTodo(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responeWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}
	var updateTodo dto.Todo
	if err := json.NewDecoder(req.Body).Decode(&updateTodo); err != nil {
		responeWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateTodo.ID = id
	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos[i] = updateTodo
			responeWithJson(writer, http.StatusOK, updateTodo)
			return
		}
	}
	responeWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}
func DeleteTodo(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responeWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid todo id"})
		return
	}
	for i, todo := range data.Todos {
		if todo.ID == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			responeWithJson(writer, http.StatusOK, map[string]string{"message": "Todo was deleted"})
			return
		}
	}
	responeWithJson(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func responeWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(todos []dto.Todo) int {
	var maxId int
	for _, todo := range todos {
		if todo.ID > maxId {
			maxId = todo.ID
		}
	}
	return maxId + 1
}
