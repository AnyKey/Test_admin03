package handlers
//
//import (
//	"awesomeProject/model"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"github.com/jmoiron/sqlx"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//type UserHandler struct {
//	Conn *sqlx.DB
//}
//
//func (uh UserHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	var err error
//
//	defer func() {
//		if err != nil {
//			log.Print(err)
//			_, _ = rw.Write(toJson(map[string]string{
//				"error": err.Error(),
//			}))
//		}
//	}()
//
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		return
//	}
//
//	user, err := uh.GetUser(id)
//	if err != nil {
//		return
//	}
//
//	_, _ = rw.Write(toJson(user))
//}
//
//func (uh UserHandler) GetUser(userID int) (*model.User, error) {
//	var user model.User
//	err := uh.Conn.Get(&user, "select id, author from book where id=?", userID)
//	return &user, err
//}
//
//func toJson(value interface{}) []byte {
//	bytes, err := json.Marshal(value)
//
//	if err != nil {
//		log.Printf("error while marshal json: %s", err)
//		return []byte{}
//	}
//
//	return bytes
//}
