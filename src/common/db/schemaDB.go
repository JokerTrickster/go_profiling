package db

type UserDTO struct{
	Name string `bson:"name"`
	Age int64 `bson:"age"`
}