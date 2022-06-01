package Routes

import (
	"github.com/aman98072/crudAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/createUser", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/getAllUsers", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/getUserById/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/updateUserById/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUserById/{id}", controllers.DeleteUser).Methods("DELETE")
}
