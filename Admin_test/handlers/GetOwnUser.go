package handlers

import (
	"awesomeProject3/model"
	"awesomeProject3/repository"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type UserHandler struct {
	Conn   *sqlx.DB
	Logger *zap.SugaredLogger
	Repo   repository.Repository
}

func (uh UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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
	var post model.Userlog
	err = json.Unmarshal(bytes, &post)
	if err != nil {
		uh.Logger.Error("Error with data")
		return
	}
	uu, err := uuid.NewV4()
	if err != nil {
		uh.Logger.Error("Error with token generate")
		return
	}
	var token string
	token = uu.String()
	tokenTest, err := uh.GetOwnUser(post.Login, post.Password, token)
	if err != nil {
		return
	}
	_ = WriteJsonToResponse(writer, interface{}(tokenTest))

	writer.WriteHeader(http.StatusOK)
}

func (uh UserHandler) GetOwnUser(login string, password string, token string) ([]model.TokenTest, error) {
	return uh.Repo.GetOwnUser(login, password, token)
}
