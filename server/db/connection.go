package db

import (
	"fmt"
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func DialMongo() *mgo.Session {
	url := fmt.Sprintf("mongodb://%s:%s@ds233278.mlab.com:33278/%s", os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASSWORD"), os.Getenv("MONGO_DB_NAME"))
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func GetCollection(collectionName string) *mgo.Collection {
	session := DialMongo()
	return session.DB(os.Getenv("MONGO_DB_NAME")).C(collectionName)
}
