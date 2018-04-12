package main

import (
	"fmt"
	"time"

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
	TimeStamp   time.Time     `bson:"timeStamp"`
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
	soup1 := Soup{
		ID:          bson.NewObjectId(),
		Name:        "ajiaco",
		Origin:      "Colombia",
		Spicy:       false,
		Ingredients: []string{"beef", "pork", "chicken", "vegetables", "starchy roots"},
		TimeStamp:   time.Now(),
	}

	//CREATE a second soup,
	soup2 := Soup{
		ID:          bson.NewObjectId(),
		Name:        "gumbo",
		Origin:      "USA",
		Spicy:       false,
		Ingredients: []string{"shrimp", "crab stock", "andouille sausage"},
		TimeStamp:   time.Now(),
	}
	//INSERT soups
	if err := c.Insert(soup1, soup2); err != nil {
		panic(err)
	}

	//FIND one by name
	queryResultByName := Soup{}
	c.Find(bson.M{"name": "ajiaco"}).One(&queryResultByName)
	fmt.Println(queryResultByName)

	//FIND by ID
	//	queryResultByID := Soup{}
	//	c.Find(bson.M{"_id": bson.ObjectIdHex("IDHERE")}).One(&queryResultByID)

	//UPDATE
	//	if err := c.Update(bson.M{"name": "ajiaco"}, bson.M{"spicy": true}); err != nil {
	//		fmt.Println("error while updating")
	//	}
	err = c.Update(bson.M{"name": "ajiaco"}, bson.M{"$set": bson.M{"spicy": true}})
	if err != nil {
		fmt.Println("error while updating")
	}

	//DELETE
	//	err = c.Remove(bson.M{"name": "gumbo"})
	//	if err != nil {
	//		fmt.Println("error removing document")
	//	}

	//COUNT
	num, err := c.Find(nil).Count()
	fmt.Println(num)

}
