package main

import (
	"encoding/json"
	//"fmt"
	"strconv"
	"time"
	"strings"
	"log"
	"os"
	//"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
)

var (
	currentPostID int
	//currentUserID int
)

// RedisConnect connects to a default redis server at port 6379
func RedisConnect() redis.Conn {
	//c, err := redis.Dial("tcp", ":6379")
	port := os.Getenv("PORT_REDIS")
	c, err := redis.Dial("tcp", ":"+ port)
	//log.Println("Redis DB Port:",port)
	HandleError(err)
	return c
}

func FindAll() Students {

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "student:*")
	HandleError(err)
	var students Students
		for _, k := range keys.([]interface{}) {
		var student Student

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &student); err != nil {
			panic(err)
		}
		//fmt.Println("---FindAll student", student)
		students = append(students, student)
		//fmt.Println("-------FindAll students",students)
	}
	//var id int64
	//id :="42"
	//reply, err := c.Do("DEL", "student:"+strconv.Itoa(id))
	//reply, err := c.Do("DEL", "student:"+id)
	//HandleError(err)

	//if reply.(int) != 1 {
	//	log.Println("No student id removed")
	//} else {
	//	log.Println("42 student id removed")
	//}
	//DeletePost(id) 
	return students
}


// CreateStudentRedis
func db_redis_CreateStudent(p Student, q *gin.Context) {

	max:= find_max_id()
	//log.Println("Logging as of",time.Now()	)
	log.Println("The max student id value is : ", max)
	
	max++
	log.Println("The next student id value is : ", max)
	//currentUserID++

	p.StudentID = max
	p.Timestamp = time.Now()
	//db_elastic_CreateStudent(p, q)
	c := RedisConnect()
	log.Println("Connected to Redis DB")
	//find_redis_keys()
	defer c.Close()

	b, err := json.Marshal(p)
	
	HandleError(err)
    
	// Save JSON blob to Redis
	//reply, err := c.Do("SET", "student:"+strconv.Itoa(p.StudentID), b)
	c.Do("SET", "student:"+strconv.Itoa(p.StudentID), b)
	log.Println("Student data saved in redis db",string(b))
	HandleError(err)

	rabbitmqSend(p, q)
	//ConsumeStudents(p, q)

}

func find_max_id() int {
	c := RedisConnect()
	defer c.Close()
	keys, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
    // handle error
	}
var a [50]int
i := 0
for _, key := range keys {
	s := strings.Split(key, ":")
    //1, 2 := s[0], s[1]
	//for i := 0; i < 10; i++ {
		a[i], err = strconv.Atoi(s[1])
		//a[i]= s[1]
	i++
	//fmt.Println(a[i])}

}
max := a[0]
for _, value := range a{
	if value > max {
			max = value 
	}
}

//fmt.Println("The biggest value is : ", max)
return max

}

func DelStudent(id int) {
	db := RedisConnect()
	defer db.Close()

	//reply, err2:=db.Do("DEL", "student:"+strconv.Itoa(id))
	db.Do("DEL", "student:"+strconv.Itoa(id))
	//HandleError(err2)
	log.Println("Student deleted:", id)
	   // if err2 != nil {
	//	log.Println("Student not deleted", id)
	//	}
	//	if err2 == nil {
	//		log.Println("Student deleted", id)
	//	}
	


	
}