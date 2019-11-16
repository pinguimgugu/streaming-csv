package repository

import (
	"fmt"

	"github.com/pinguimgugu/streaming-csv/domain/entity"
	"github.com/pinguimgugu/streaming-csv/infrastructure/mongodb"
	"gopkg.in/mgo.v2/bson"
)

// User repository using mongodb
type User struct {
	Connection *mongodb.Connection
}

// StreamingAllUsers method
func (u *User) StreamingAllUsers() chan *entity.User {
	chanUsers := make(chan *entity.User)

	coll := u.Connection.Db.C("users")
	var results []*entity.User
	coll.Find(bson.M{}).All(&results)

	go func(uChan chan *entity.User) {
		defer close(uChan)
		for _, data := range results {
			fmt.Println(data)
			uChan <- data
		}
	}(chanUsers)

	return chanUsers
}
