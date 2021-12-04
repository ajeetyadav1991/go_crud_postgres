package internal

import (
	"github.com/gorilla/mux"
)

//create a new router
// `Router` starts with capital character, so, it is exported and used in main.go
func Router() *mux.Router {
	//The Router function will handle all the endpoints and respective handlers.

	router := mux.NewRouter()

	router.HandleFunc("/", HomePageHandler).Methods("GET") // HomePage

	router.HandleFunc("/api/getuser/{id}", GetUserHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getuserlist", GetAllUsersHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/createnewuser", CreateUserHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteuser/{id}", DeleteUserHandler).Methods("DELETE", "OPTIONS")

	//for API endpoints which perform partial updates on a resource,  we use the PATCH rather than PUT
	// PUT is intended for replacing a resource in full
	router.HandleFunc("/api/updateuser/{id}", UpdateUserHandler).Methods("PATCH", "OPTIONS")

	return router
}
