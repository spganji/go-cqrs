
package main

import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"io/ioutil"
	"log"
	"net/http"
//	"strconv"
//	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"

//	mux "github.com/julienschmidt/httprouter"
)

var (
	elasticClient *elastic.Client
)

const (
	elasticIndexName = "documents"
	elasticTypeName  = "document"
	)

	/*
type DocumentRequest struct {
	//StudentID int       `json:"StudentID"`
	Name   string `json:"Name"`
	Course string `json:"Course"`
	//Timestamp time.Time `json:"timestamp"`
} */


func db_elastic_CreateStudent(student Student, c *gin.Context) {
	
	var err error
	//connect to elastic search
	elasticClient, err = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	
	if err != nil {
		log.Println(err)
	}
	log.Println("Connected to Elastic DB")
	// Parse request
	/*var docs DocumentRequest
	if err := c.BindJSON(&docs); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}*/
	//fmt.Println("--------------3")
	// Insert documents in bulk
	bulk := elasticClient.
		Bulk().
		Index(elasticIndexName).
		Type(elasticTypeName)
		// for _, d := range docs {
	//ID := shortid.MustGenerate()
	//p.StudentID = currentPostID
	//p.User.ID = currentUserID
	//p.Timestamp = time.Now()
	
	log.Println("Elastic indexing...")
	doc := Student{
		StudentID: student.StudentID,
		Name:      student.Name,
		Course:    student.Course,
		Timestamp: student.Timestamp,
		//Timestamp: time.Now(),
		//Timestamp: time.Now().UTC(),		
	}

	bulk.Add(elastic.NewBulkIndexRequest().Id(doc.Name).Doc(doc))
	// }
	if _, err := bulk.Do(c.Request.Context()); err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Failed to create documents")
		return
	}
	c.Status(http.StatusOK)
	log.Println("Student data saved in elastic db", doc)
}