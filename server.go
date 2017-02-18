package main

import(
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql" //database driver
)

//DB Connection Details
const(
	DBHost = "127.0.0.1"
	DBPort = "3306"
	DBUser = "root"
	DBPass = "testing123"
	DBDbase = "contakx"
)

//Database identifier based on gorm.DB
var AppDB *gorm.DB

func init(){
	var err error

	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBDbase)

	AppDB, err = gorm.Open("mysql", dbConnection)

	if err != nil {
		log.Println("Failed to connect to database: ", err.Error())
	}else {
		log.Println("DB Connection: connected to the database successfully")
	}

	AppDB.Create(&Product{Code:"L1212", Price: 1000})
	AppDB.Create(&Contact{FirstName:"Jorge", LastName:"Astorga"})

	//defer AppDB.Close()

	//Migrate Schema
	AppDB.AutoMigrate(&Contact{}, &Product{})

	//Testing the creation of a contact
	newContact := &Contact{FirstName:"Jorge", LastName:"Astorga"}
	newContact.create()
}

func main(){

	//Testing to add a contact
	AppDB.Create(&Contact{FirstName:"Oralge", LastName:"Orale"})


	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/contact/", handleRequest)
	server.ListenAndServe()

}



func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error
	log.Println("testing")
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
 *
 *
 */
func handleGet(w http.ResponseWriter, r *http.Request) (err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	contact, err := retrieve(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&contact, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var contact Contact
	json.Unmarshal(body, &contact)
	err = contact.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

/**
 *
 */
func handlePut(w http.ResponseWriter, r *http.Request)(err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	contact, err := retrieve(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &contact)
	err = contact.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}


/**
 *
 */
func handleDelete(w http.ResponseWriter, r *http.Request) (err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	contact, err := retrieve(id)
	if err != nil {
		return
	}

	err = contact.delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}