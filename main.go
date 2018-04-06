package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Soup export
type Soup struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Origin      string        `bson:"origin"`
	Spicy       bool          `bson:"spicy"`
	Ingredients []string      `bson:"ingredients"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("mongodb connected")

	defer session.Close()

	//Collection
	c := session.DB("RECEPIES").C("soups")

	//CREATE a soup
	//	soup1 := Soup{
	//		ID:          bson.NewObjectId(),
	//		Name:        "ajiaco",
	//		Origin:      "Colombia",
	//		Spicy:       false,
	//		Ingredients: []string{"beef", "pork", "chicken", "vegetables", "starchy roots"},
	//	}
	//
	//	//CREATE a second soup,
	//	soup2 := Soup{
	//		ID:          bson.NewObjectId(),
	//		Name:        "gumbo",
	//		Origin:      "USA",
	//		Spicy:       false,
	//		Ingredients: []string{"shrimp", "crab stock", "andouille sausage"},
	//	}

	//INSERT soups
	//	if err := c.Insert(soup1, soup2); err != nil {
	//		panic(err)
	//	}

	//FIND one by name
	//queryResult := Soup{}
	//c.Find(bson.M{"name": "ajiaco"}).One(&queryResult)
	//fmt.Println(queryResult)

	//FIND by ID
	queryResult := Soup{}
	c.Find(bson.M{"_id": bson.ObjectIdHex("5ac770a55580180c8e12fdd5")}).One(&queryResult)
	fmt.Println(queryResult.ID)
}
