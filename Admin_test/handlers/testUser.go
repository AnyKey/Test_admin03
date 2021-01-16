package handlers

//type UserHandler struct {
//	Conn   *sqlx.DB
//	Logger *zap.SugaredLogger
//	Repo   repository.UserRepository
//}
//
//func (handler UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Access-Control-Allow-Origin", "*")
//	var err error
//	defer func() {
//		if err != nil {
//			_ = WriteJsonToResponse(writer, map[string]string{
//				"error": err.Error(),
//			})
//		}
//		writer.WriteHeader(http.StatusOK)
//	}()
//	_ = WriteJsonToResponse(writer, handler)
//}
//func(handler UserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	writer.Header().Set("Access-Control-Allow-Origin", "*")
//	var err error
//	defer func() {
//		if err != nil {
//			_ = WriteJsonToResponse(writer, map[string]string{
//				"error": err.Error(),
//			})
//		}
//		writer.WriteHeader(http.StatusOK)
//	}()
//	bytes, _ := ioutil.ReadAll(request.Body)
//	var post model.User
//	err = json.Unmarshal(bytes, &post)
//	if err != nil {
//		logger.Error("Error with data")
//		return
//	}
//
//	_ = handler.WriteJsonToResponse(writer, handler)
//	writer.WriteHeader(http.StatusOK)
//}
//func (handler UserHandler) GetUser() ([]model.User, error) {
//	return handler.Repo.GetUser()
//}
//
//func (handler UserHandler) AddUser(login string, password string) ([]model.User, error) {
//	return handler.Repo.AddUser(login, password)
//}
