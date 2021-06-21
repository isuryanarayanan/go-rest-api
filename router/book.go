package router

import (
	"github.com/gorilla/mux"
	"github.com/isuryanarayanan/go-rest-api/controllers"
)

func BookRouter(r *mux.Router) {
	r.HandleFunc("", controllers.BookView)
}
