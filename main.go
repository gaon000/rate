package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReturnAll(client *mongo.Client, filter bson.M, db_name string, collection_name string) []*Officer {
	var officers []*Officer
	collection := client.Database(db_name).Collection(collection_name)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("err on finding document", err)
	}
	for cur.Next(context.TODO()) {
		var officer Officer
		err = cur.Decode(&officer)
		if err != nil {
			log.Fatal("err on decoding doc", err)
		}
		officers = append(officers, &officer)
	}
	return officers
}

type Officer struct {
	Occupation string "json:'occupation'"
	Applicant  string "json:'applicant'"
	Selected   string "json:'selected'"
	Rate       string "json:'rate'"
}

/*type Trainer struct {
    Name string
    Age  int
    City string
}
*/

var db = make(map[string]string)

func setupRouter(mc *mongo.Client) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	//r.LoadHTMLGlob("practice/*")

	r.GET("/five", func(c *gin.Context) {
		var a []string
		officers := ReturnAll(mc, bson.M{}, "officer", "five")
		for _, officer := range officers {
			a = append(a, officer.Occupation, officer.Applicant, officer.Selected, officer.Rate)
		}
		c.JSON(200, a)
	})

	r.GET("seven", func(c *gin.Context) {
		var a []string
		officers := ReturnAll(mc, bson.M{}, "officer", "seven")
		for _, officer := range officers {
			a = append(a, officer.Occupation, officer.Applicant, officer.Selected, officer.Rate)
		}
		c.JSON(200, a)
	})

	r.GET("/nine", func(c *gin.Context) {
		var a []string
		officers := ReturnAll(mc, bson.M{}, "officer", "nine")
		for _, officer := range officers {
			a = append(a, officer.Occupation, officer.Applicant, officer.Selected, officer.Rate)
		}
		c.JSON(200, a)
	})

	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.GET("/officer", func(c *gin.Context) {
		a := [3]string{"five", "seven", "nine"}
		c.JSON(200, a)
	})

	// Ping test
	r.GET("/", func(c *gin.Context) {

	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	c := GetClient()
	// err := c.Ping(context.Background(), readpref.Primary())
	// if err != nil {
	// log.Fatal("couldn't connect to db", err)
	// } else {
	// log.Println("connected")
	// }
	// findOptions := options.Find()
	// findOptions.SetLimit(2)
	// collection = client.Database("test").Collection("trainers")
	// }
	// collection_nine := client.Database("officer").Collection("nine")
	// collection_seven := client.Database("officer").Collection("seven")
	// collection_five := client.Database("officer").Collection("five")
	// cur, err := collection_nine.Find(context.TODO(), bson.D{{}}, findOptions)
	// if err != nil {
	// log.Fatal(err)
	// }
	// for cur.Next(context.TODO()) {
	// var elem Officer
	// err := cur.Decode(&elem)
	// if err != nil {
	// log.Fatal(err)
	// }
	// result = append(result, &elem)
	// }
	// cur.Close(context.TODO())
	// fmt.Printf("%+v\n", result)

	r := setupRouter(c)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
