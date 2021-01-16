package handlers

import (
	"awesomeProject3/model"
	"awesomeProject3/repository"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type AUserHandler struct {
	Conn   *sqlx.DB
	Logger *zap.SugaredLogger
	Repo   repository.Repository
}

func (auh AUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	defer func() {
		if err != nil {
			_ = WriteJsonToResponse(writer, map[string]string{
				"error": err.Error(),
			})
		}
		writer.WriteHeader(http.StatusOK)
	}()
	bytes, _ := ioutil.ReadAll(request.Body)
	var post model.User
	err = json.Unmarshal(bytes, &post)
	if err != nil {
		auh.Logger.Error("Error with data")
		return
	}

	err = auh.AddUser(post.Login, post.Password)
	if err != nil {
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (auh AUserHandler) AddUser(login string, password string) (error) {
	return auh.Repo.AddUser(login, password)
}
