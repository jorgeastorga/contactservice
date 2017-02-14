package main

import(
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"log"
)


func main(){
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