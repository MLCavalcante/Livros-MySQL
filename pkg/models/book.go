package models

import (
	"github.com/jinzhu/gorm" 
	"github.com/MLCavalcante/sgb-mysql/pkg/config" //ajuda a se conectar com a database
)

var db *gorm.DB

type Book struct { //como se fosse um modelo de estrutura pra as coisas que serão guardadas na database
   gorm.Model
   Name string `gorm: ""json:"name"`
   Author string `json:"author"`     //editar database
   Publication string `json:"publication`  
 
}

func init(){  // funcção que ajuda a iniciar a database 
	config.Connect()
	db = config.GetDB()  //db que pegamos da config file
	db.AutoMigrate(&Book{}) // cria/atualiza automaticamente as tabelas do banco de dados, o gorm analisa a struct e mantéma a tabela em sincronia


}

//o orm gorm reduz a conversa (escrita de querys) com a database em pequenas funções 

                                                                                     
  
func (b *Book) CreateBook() *Book { 
	db.NewRecord(b)   // b é algo que recebos de type Book e o que retornamos tb
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
  var Books []Book 
  db.Find(&Books) // busca todos os registros associados à estrutura Book no banco de dados e preenche o slice Books com os resultados.
   return Books  //a função retorna Books, que é um slice do tipo []Book, do banco de dados 
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book // cria a var vazia getBook que é do tipo Book
	db := db.Where("ID=?", Id).Find(&getBook)// "na condição de ID ser igual id, execute uma consulta e preencha a var com o resultado"
    return &getBook, db

}

func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?", Id).Delete(book) // ""na condição de ID ser igual id, delete"
	return book
}

// O update será feio pela combinação desses 3 métodos acharemos o livro pelo id deletaremos a data e criaremos outra