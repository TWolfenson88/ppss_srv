package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//TODO: Should create config reader for database
//also maybe should create separation between graph_editor and ppss databases??

//DBconf struct is a database configuration
type DBconf struct {
	Host   string
	Port   uint16
	User   string
	Pass   string
	DBname string
}

//Dbase hadle a connections to DB
type Dbase struct {
	*sql.DB
}

//Global configuration of database
var DBcfg = &DBconf{}

//InitDB func get config data and initialize database
func InitDB(cfg *DBconf) (*Dbase, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	log.Println("Base initialezed!")
	return &Dbase{db}, nil
}
