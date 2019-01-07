package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"           //MySQL database driver
	"github.com/jinzhu/gorm"                     //TODO update package diagram in Lucidcharts to include this package
	_ "github.com/jinzhu/gorm/dialects/postgres" //PostgreSQL database driver
	"io/ioutil"
	"log"
	_ "log"
	"net/http"
)


//Database identifier based on gorm.DB
var AppDB *gorm.DB


/**
 * Init (function) -- Package initialization
 * Defines the connection string for the mysql connection, opens the connection, and migrates to
 * a new schema (based on objects)
 */
func init() {

	newContact := &Contact{FirstName: "Adriana", LastName: "Astorga"}
	newContact.create()
}

/**
 * main - (function)
 * Starts the contact (micro)service
 */
func main(){

	/**
	* TODO: AWS-specific need to investigate why not using the IP address of the instance and also investigate
	* if I should bind to the private IP as opposed to the public IP or 0.0.0.0:8080
	**/
	server := http.Server{
		//Addr: "127.0.0.1:8080",
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/contact/", handleRequest)
	server.ListenAndServe()
}


/**
 * HandleRequest
 * Function that looks at the request method (GET, POST, PUT, AND DELETE) and dispatches code execution to the
 * appropriate function.
 */

func handleRequest(w http.ResponseWriter, r *http.Request){

	var err error

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

/**
 * HandleGet (function) -- READ
 * Reads/Retrieves a specific contact from the data store
 */
func handleGet(w http.ResponseWriter, r *http.Request) (err error){

	contact := &Contact{FirstName:"Lucca", LastName:"Astorga"}

	log.Println("Retrived a contact")

	output, err := json.MarshalIndent(&contact, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)


	/************** Adding some code to call an API ************/
	fmt.Print("/nCalling an API/n")

	response, err := http.Get("https://httpbin.org/ip") //call the API
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}


	/************** Adding some code to call an API ************/

	jsonData := map[string]string{"firstname":"Nic", "lastname":"Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("Terminating the application....")

	return
}

/**
 * HandlePost (function) -- WRITE/CREATE
 * Creates a contact in the datastore
 */

func handlePost(w http.ResponseWriter, r *http.Request) (err error){
	//Read the request data & json doc
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	//Create contact
	var contact Contact
	json.Unmarshal(body, &contact)

	contact.create()

	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

/**
 * HandlePut (function) - UPDATE
 * Updates an existing contact in the datastore
 *
 */
func handlePut(w http.ResponseWriter, r *http.Request)(err error){


	//retrieve the contact from the datastore
	contact := &Contact{FirstName:"Update", LastName:"Astorga"}


	//retrieve the updated info from request (json doc)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &contact)

	//update the contact in the datastore
	contact.update()

	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}


/**
 * HandleDelete (function) - DELETE
 * Deletes a contact from the datastore
 */
func handleDelete(w http.ResponseWriter, r *http.Request) (err error){

	if err != nil {
		return
	}

	//Retrieve the contact from the datastore (if it exists)
	contact := &Contact{FirstName:"Update", LastName:"Astorga"}

	if err != nil {
		return
	}

	//Delete the contact from the datastore
	 contact.delete()

	w.WriteHeader(200)
	return
}