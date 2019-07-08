package main

import (
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//"github.com/olivere/elastic"

	//mux "github.com/julienschmidt/httprouter"
)

/*var (
	elasticClient *elastic.Client
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
)

type DocumentRequest struct {
	//StudentID int       `json:"StudentID"`
	Name   string `json:"Name"`
	Course string `json:"Course"`
	//Timestamp time.Time `json:"timestamp"`
}*/

// Logger logs the method, URI, header, and the dispatch time of the request.

func init(){
    //Logging...
	f, _ := os.Create("go-redis-elastic.log")
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter) 
	log.Println("Logging as of",time.Now()	)
	
	//env file setup	
	log.Println("Loading go.env file")
	err2 := godotenv.Load("go.env")
	if err2 != nil {
	log.Fatal("Error loading .env file")
	}
	
	}
/*
func Logger(r *http.Request) {
	//start := time.Now()
	/*log.Printf(
		"%s\t%s\t%q\t%s",
		r.Method,
		r.RequestURI,
		r.Header,
		time.Since(start),
	)
} */


// GetStudents handler queries the students 
func GetStudents(c *gin.Context) {
	//httpMethod := c.Request.Method
	//sp:= c.HandlerName()
	
	//httpMethod2 := c.handl
	//log.Println("httpMethod:",httpMethod, sp)
	w := c.Writer
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//elasticdbconn()
	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}


// func CreateStudent(w http.ResponseWriter, r *http.Request, _ mux.Params) {
func CreateStudent(c *gin.Context) {
	//w http.ResponseWriter, r *http.Request, _ mux.Params,
	r := c.Request
	w := c.Writer

	//Logger(r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Post struct (should this be a pointer?)
	var student Student
	if err := json.Unmarshal(body, &student); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
	//redis db method & elastic db method
	db_redis_CreateStudent(student,c)
	//rabbitmqSend(student, c)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func StudentDelete(c *gin.Context) {
	//r := c.Request
	//w := c.Writer
	ps:= c.Params
	//log.Println("1-------", ps)
	db := RedisConnect()
	defer db.Close()
	//log.Println("1-------")
	id, err := strconv.Atoi(ps.ByName("studentid"))
	//log.Println("2-------", id)
	HandleError(err)
	DelStudent(id)
		
}


