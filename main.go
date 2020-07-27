package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"todos-backend/controllers"
	"todos-backend/models"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "tododb"
)

func getLists(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("get worked")
	//	payload, err := json.Marshal(q)
	//	if err != nil {
	//		panic(err)
	//	}
	//	w.Write(payload)
}

func main() {
	//TODO make a createDB method to check that the DB exists
	//creates DB if no DB with the same name exists
	//	createDB()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, os.Getenv("LOGNAME"), os.Getenv("LOGNAME"), dbname)
	list, err := models.NewListService(psqlInfo)
	must(err)
	defer list.Close()
	list.AutoMigrate()
	listsC := controllers.NewLists(list)
	router := mux.NewRouter()

	l, err := list.ByID(1)
	if err != nil {
		panic(err)
	}
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.StrictSlash(true)
	fmt.Println("first list is ", l)
	router.HandleFunc("/api/lists", getLists).Methods("GET")
	router.HandleFunc("/api/lists/create", listsC.Create).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

//func createDB() {
//	statement := `SELECT 1 FROM pg_database WHERE datname='tododb';`
//	row := db.QueryRow(statement)
//	var exists bool
//	err := row.Scan(&exists)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("exists is ", exists)
//
//}
