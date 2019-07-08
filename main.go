package main

import (
	"log"
"os"
//"fmt"
//"time"
//"io"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	//"github.com/olivere/elastic"
	//"github.com/teris-io/shortid"
)



func main() {  
	var err error

	// Start HTTP server
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	r.POST("/createstudent", CreateStudent)
	r.GET("/getstudents", GetStudents)
	r.GET("/consumestudents", ConsumeStudents)
	r.GET("/search", SearchAPI)
	r.POST("/delstudent/:studentid", StudentDelete)
	
	port := os.Getenv("PORT")
  	//if err = r.Run(":8080"); err != nil {
		if err = r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
	
}

// HandleError handles error.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

