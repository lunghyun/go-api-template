package student

import "github.com/gorilla/mux"

// RegisterRoutes student 라우팅 등록
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/students", GetListHandler).Methods("GET")
	r.HandleFunc("/students/{id:[0-9]+}", GetHandler).Methods("GET")
	r.HandleFunc("/students", PostHandler).Methods("POST")
	r.HandleFunc("/students/{id:[0-9]+}", DeleteHandler).Methods("DELETE")
}
