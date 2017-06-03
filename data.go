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
func (c *Contact)create() {

	log.Println("Created a contact")
}

/**
* Retrieve (function) -
* Function used to retrieve a Contact from the database
*
*/
func retrieve(id int)  {
	log.Println("Retrived a contact")

}

/**
* Update (function) -
* Function used to update a Contact from the database
*
*/
//TODO: update the update method to use Gorm
func (c *Contact) update() {
	log.Println("Updated the contact")
}

/**
* Delete (function) - delete contact
* Function used to delete a Contact from the db
*
*/
//TODO: update the delete method to use Gorm
func (c *Contact)delete() {
	log.Println("Deleted a contact")
}

