package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bikashsaud/mongodb_api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://bikashsaud:Djangoapp514@cluster0.icctrwy.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// Most important
var collection *mongo.Collection

// connect to MongoDB
func init() { //only first time to initialize
	// client connection
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to MongoDB
	client, err := mongo.connect(context.TODO(), clientOption) //context.TODO and context.Background two ways...
	if err != nil {
		// panic("error", err)
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection established.")
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("collection instance is ready", collection)
}

// mongoDB helpers ---> make separate files is better practice

func addMovie(movie models.Netflix) {
	add, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted data to bd", add.InsertedId)
}

// update a movie
func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	update, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modify count:", update.ModifiedCount)
}

func deleteMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	delete, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("delete object:", delete)
}

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteall, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted all movies: ", deleteall.DeleteCount)
	return deleteall.DeleteCount
}

func getMovies() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// controller file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode") //application/json
	allMovies := getMovies()

	json.NewEncoder(w).Encode(allMovies)
	fmt.Println("all movies: ", allMovies)
	return
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	addMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// get ad from params

	params := mux.Vars(r)
	updateMovie(params["id"])

	fmt.Println("Update Movie: ", params["id"])
	json.NewEncoder(w).Encode(params)
	return

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
