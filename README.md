This is a simple project for introducing the basics of a web app in Go. The focus will be routing http requests, saving and retrieving data from a Mongodb instance.

MONGO DB

We are using the mgo package (more explanation later).

First of all we have to start with a session instance by calling the Dial function form the mgo package and passing in the url string of Mongodb. The port defaults to port 27017.


Struct vs Maps to Find()

The Find method receives an empty interface and returns a pointer to a Query.  The value of the empty interface must be a struct or a map.  The mgo/bson library has a bson map(bson.M) which is less verbose than marshalling a stuct to bson and passing to the find method. 