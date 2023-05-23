package main // criar um server que tb vai definir localhost
             // essa página basicamente aponta para onde as rotas estão 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/MLCavalcante/sgb-mysql/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)	 


func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}