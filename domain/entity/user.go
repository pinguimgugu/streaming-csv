package entity

//User struct
type User struct {
	ID   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}
