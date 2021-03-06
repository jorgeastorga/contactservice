package main

import (
	_"github.com/lib/pq"
	"log"
	"github.com/jinzhu/gorm"
)


/**
* Product (structure)
* Product data model and definition
*
*/
type Contact struct {
	gorm.Model // declaring the model (based on Gorm)
	//Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode int `json:"zip"`
	Company string `json:"company"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

/**
 * Create (function) -
 * Function used to create a contact in the database
 */
func (c *Contact)create() (err error){

	//Create the contact
	err = AppDB.Create(&Contact{FirstName:c.FirstName, LastName: c.LastName}).Error

	if err != nil{
		log.Println("Error Creating the Contact: ", err.Error())
	}

	return err
}

/**
* Retrieve (function) -
* Function used to retrieve a Contact from the database
*
*/
func retrieve(id int) (contact Contact, err error) {

	AppDB.First(&contact, id)
	return contact, err
}


/**
* Update (function) -
* Function used to update a Contact from the database
*
*/
//TODO: update the update method to use Gorm
func (c *Contact) update() (err error){
	/*_, err = Db. Exec("update contacts set firstname = $2, lastname = $3, address1 = $3 where id = $1",
		c.Id, c.FirstName, c.LastName, c.Address1)*/
	return err
}

/**
* Delete (function) - delete contact
* Function used to delete a Contact from the db
*
*/
//TODO: update the delete method to use Gorm
func (c *Contact)delete() (err error){
	//_, err = Db. Exec("delete from contacts where id = $1", c.Id)
	return
}

