package config

// a função dessa file é retornar essa var db, que vai ajudar as outras files a interagir com o db

import ( 
	"github.com/jinzhu/gorm" //orm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

// essa func serve pra abrir a concexão com a database

func Connect() {
	d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
