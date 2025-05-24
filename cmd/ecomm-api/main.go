package main

import (
	"github.com/bloodgroup-cplusplus/ecomm/db"
)
func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database :%v",err)
	}
	defer db.Close()
	log.Prinln("Successfully connected to database")

	// do something with the databawse
}




