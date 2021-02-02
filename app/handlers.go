//Package app contain functions and methods to handle requests from app
package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"test_server/db"
)

type Hndle struct {
	dbs *db.Dbase
}

//Handler func contain router to handle all app requests
func Handler(s *mux.Router) {
	db, err := db.InitDB(db.DBcfg)
	if err != nil {
		log.Fatal(db)
	}

	hnd := &Hndle{db}

	fmt.Println("HHMMMM")

	s.HandleFunc("/get/GetAllStans", hnd.GetAllStans)
	s.HandleFunc("/get/GetStans", hnd.GetStans)
	s.HandleFunc("/get/NextStans", hnd.GetNextStans)


}

//GetAllStans method return JSON array with all stantions
func (h Hndle) GetAllStans(w http.ResponseWriter, r *http.Request) {
	fmt.Println("YAY< YOU GOT IT!")

	stnsArr, err := h.dbs.GetAllStans()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("500 - Something on server goes wrong!"))
	}

	if len(stnsArr) != 0 {
		mrshlStns, err := json.Marshal(stnsArr)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("500 - Something on server goes wrong!"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(mrshlStns)
	} else {
		w.WriteHeader(204)
	}
}

//GetStans method return JSON with filtered stations by road and/or flag
func (h Hndle) GetStans(w http.ResponseWriter, r *http.Request) {
	dorCode, _ := strconv.Atoi(r.URL.Query().Get("dorCode"))
	flag := r.URL.Query().Get("flag")

	stnsArr, err := h.dbs.GetStans(dorCode, flag)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("500 - Something on server goes wrong!"))
	}

	if len(stnsArr) != 0 {
		mrshlStns, err := json.Marshal(stnsArr)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("500 - Something on server goes wrong!"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(mrshlStns)
	} else {
		w.WriteHeader(204)
	}
}

//GetNextStans method return list of all stations around the current stan
func (h Hndle) GetNextStans(w http.ResponseWriter, r *http.Request) {
	dorCode, _ := strconv.Atoi(r.URL.Query().Get("cnsi"))

	stnsArr, err := h.dbs.GetNextStans(dorCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("500 - Something on server goes wrong!"))
	}

	if len(stnsArr) != 0 {
		mrshlStns, err := json.Marshal(stnsArr)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
			w.Write([]byte("500 - Something on server goes wrong!"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(mrshlStns)
	} else {
		w.WriteHeader(204)
	}
}
