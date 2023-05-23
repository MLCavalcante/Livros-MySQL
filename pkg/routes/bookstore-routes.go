package routes

import (
	"github.com/gorilla/mux"
	"github.com/MLCavalcante/sgb-mysql/pkg/controllers"
)

//as rotas ajudam a acessar os controllers essa função vai guardar todas as rotas  da aplicação
var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	 
}

