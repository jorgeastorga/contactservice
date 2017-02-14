package main

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init(){
	log.Println("init called")
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp ssl-mode=disable")
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}

}


type Contact struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode int `json:"zip"`
	Company string `json:"compnay"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

/**
 * Function used to create  a contact in the db
 */
func (c *Contact)create() (err error){
	statement := "insert into contacts (firstname, lastname, address1) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(c.FirstName, c.LastName, c.Address1).Scan(&c.Id)
	return err
}

/**
* Function used to retrieve a Contact from the db
*
*/
func retrieve(id int) (contact Contact, err error) {
	contact = Contact{}
	err = Db.QueryRow("select id, firstName, lastName, address1 from contacts where id = $1",
		id).Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Address1)
	return contact, err

}


/**
* Function used to update a Contact from the db
*
*
*/
func (c *Contact) update() (err error){
	_, err = Db. Exec("update contacts set firstname = $2, lastname = $3, address1 = $3 where id = $1",
		c.Id, c.FirstName, c.LastName, c.Address1)
	return err
}

/**
* Function used to delete a Contact from the db
*
 */
func (c *Contact)delete() (err error){
	_, err = Db. Exec("delete from contacts where id = $1", c.Id)
	return
}

