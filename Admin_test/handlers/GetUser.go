package handlers

import (
	"awesomeProject3/model"
	"awesomeProject3/repository"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"net/http"
)

type GUserHandler struct {
	Conn   *sqlx.DB
	Logger *zap.SugaredLogger
	Repo   repository.Repository
}

func (guh GUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	token := request.Header.Get("x-token")
	//vars := mux.Vars(request)

	var err error
	defer func() {
		if err != nil && sql.ErrNoRows != err {
			writer.WriteHeader(http.StatusBadGateway)
			_ = WriteJsonToResponse(writer, map[string]string{
				"error": err.Error(),
			})
		}

	}()
	auth, err := guh.CheckToken(token)
	if err != nil {
		return
	}
	if auth == true {
		users, err := guh.GetUser()
		if err != nil {
			return
		}
		writer.WriteHeader(http.StatusOK)
		_ = WriteJsonToResponse(writer, interface{}(users))
	} else {
		errorToken := model.ErrorTest{ErrorT: "Error authorization token!"}
		if err != nil {
			return
		}
		writer.WriteHeader(http.StatusUnauthorized)
		_ = WriteJsonToResponse(writer, interface{}(errorToken.ErrorT))

	}
}

func (guh GUserHandler) GetUser() ([]model.User, error) {
	return guh.Repo.GetUser()
}

func (guh GUserHandler) CheckToken(token string) (model.Authorization, error) {
	return guh.Repo.CheckToken(token)
}
