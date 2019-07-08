package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	//"io"
    //"os"
	//"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	//"github.com/teris-io/shortid"
)

/*type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	//Timestamp time.Time `json:"timestamp"`
	Content   string    `json:"content"`
}*/

type DocumentResponse struct {
	StudentID        int       `json:"StudentID"`
	Name      string    `json:"Name"`
	Course    string    `json:"Course"`
	Timestamp time.Time `json:"Timestamp"`
}

type SearchResponse struct {
	//Time      string             `json:"time"`
	//Hits      string             `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}

func SearchAPI(c *gin.Context) {
	var err error
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	// Parse request
	query := c.Query("query")
	if query == "" {
		errorResponse(c, http.StatusBadRequest, "Query not specified")
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(c.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(c.Query("take")); err == nil {
		take = i
	}
	// Perform search
	esQuery := elastic.NewMultiMatchQuery(query, "Name", "Course").
		Fuzziness("2").
		MinimumShouldMatch("2")
	result, err := elasticClient.Search().
		Index(elasticIndexName).
		Query(esQuery).
		From(skip).Size(take).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		errorResponse(c, http.StatusInternalServerError, "Something went wrong")
		return
	}
	res := SearchResponse{
		//Time: fmt.Sprintf("%d", result.TookInMillis),
		//Hits: fmt.Sprintf("%d", result.Hits.TotalHits),
	}
	// Transform search results before returning them
	docs := make([]DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc DocumentResponse
		json.Unmarshal(hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	c.JSON(http.StatusOK, res)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}