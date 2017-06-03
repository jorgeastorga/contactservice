package main

import(
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"log"
	"fmt"
	"github.com/jinzhu/gorm" //TODO update package diagram in Lucidcharts to include this package
	_"github.com/go-sql-driver/mysql" //MySQL database driver
	_"github.com/jinzhu/gorm/dialects/postgres" //PostgreSQL database driver
)

//DB Connection Details
//TODO: Find a way to make this information more secure and/or read from a config file
const(
	DBHost = "127.0.0.1"
	DBPort = "3306"
	DBUser = "root"
	DBPass = "testing123"
	DBDbase = "contakx"
)

//Database identifier based on gorm.DB
var AppDB *gorm.DB


/**
 * Init (function) -- Package initialization
 * Defines the connection string for the mysql connection, opens the connection, and migrates to
 * a new schema (based on objects)
 */
func init() {
	var err error

	//Setup connection string
	//TODO remove MYSQL connectin string once PostgreSQL works
	/*dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		DBUser,
		DBPass,
		DBHost,
		DBPort,
		DBDbase)*/

	dbConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		DBHost,
		DBUser,
		DBDbase,
		"disable",
		DBPass)

	//Open DB - Test Connection
	AppDB, err = gorm.Open("postgres", dbConnection)

	//TODO Remove the reference to MYSQL connection open
	//AppDB, err = gorm.Open("mysql", dbConnection)

	if err != nil {
		log.Println("Failed to connect to database: ", err.Error())
	} else {
		log.Println("DB Connection: connected to the database successfully")
	}

	//defer AppDB.Close()

	//Migrate Schema
	AppDB.AutoMigrate(&Contact{})

	//TODO: Remove this code
	//Testing the creation of a contact
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

	//TODO: understand this code in more depth

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
	err = contact.create()

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
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	//retrieve the contact from the datastore
	contact, err := retrieve(id)
	if err != nil {
		return
	}

	//retrieve the updated info from request (json doc)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &contact)

	//update the contact in the datastore
	err = contact.update()
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

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	//Retrieve the contact from the datastore (if it exists)
	contact, err := retrieve(id)
	if err != nil {
		return
	}

	//Delete the contact from the datastore
	err = contact.delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}