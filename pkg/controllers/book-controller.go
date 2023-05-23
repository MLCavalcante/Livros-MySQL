package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	
	"github.com/MLCavalcante/sgb-mysql/pkg/models"
	"github.com/MLCavalcante/sgb-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	NewBooks:= models.GetAllBooks()
	res, _ := json.Marshal(NewBooks) // converte em json oq vemos na database
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // manda a res p insomnia/frontend (no caso um json com os livros que achamos na database)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // vamos acessar o r que é o nosso request 
	bookId := vars["bookId"] // acessar o book id dentro do request 
	ID, err := strconv.ParseInt(bookId,0,0) //converte p int 
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // manda a res p insomnia/frontend (no caso um json com o livro by id)
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := models.Book{} // recebemos json
	utils.ParseBody(*r , CreateBook) // passamos o r para algo que a database/golang entenda 
	b:= CreateBook.CreateBook() 
	res, _ := json.Marshal(b) // passamos para json e mandamos para o "user" insomnia/frontend
    w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // vamos acessar o r que é o nosso request 
	bookId := vars["bookId"] // acessar o book id dentro do request 
	ID, err := strconv.ParseInt(bookId,0,0) //converte p int 
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := models.Book{} // recebemos json     
	utils.ParseBody(*r , updateBook) // passamos o r para algo que o golang/database entenda 
	vars := mux.Vars(r)
	bookId := vars["bookId"] // acessar o book id dentro do request 
	ID, err := strconv.ParseInt(bookId,0,0) //converte p int 
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID) //pegamos book details de getbookbyid
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
