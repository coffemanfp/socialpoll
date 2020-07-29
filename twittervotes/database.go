package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Session
var databaseHost string = "localhost"

func dialdb() (err error) {
	log.Println("dialing mongodb: ", databaseHost)
	db, err = mgo.Dial(databaseHost)
	return
}

func closedb() {
	db.Close()
	log.Println("closed database connection")
}
