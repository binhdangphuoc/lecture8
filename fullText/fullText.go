package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func es()  (string,time.Duration){
	mes := "no"
	esStart := time.Now()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	//log.Println(res)
	var (
		r  map[string]interface{}
		//wg sync.WaitGroup
		buf bytes.Buffer
	)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"body": "Stuning",
			},
		},
	}
	err = json.NewEncoder(&buf).Encode(query)
	if err!=nil{
		fmt.Println("Error encode query")
	}
	res,err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("review"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)
	if res.Status()=="200 OK"{
		mes = "yes"
	}
	esEnd := time.Now()
	elapse := esEnd.Sub(esStart)
	// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}

	//log.Println(strings.Repeat("=", 37))

	return mes,elapse
}
func mysql() (string,time.Duration){
	mes := "no"
	sqlStart := time.Now()
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err!=nil{
		logrus.Error("Can not connect mysql")
	}
	res,err := db.Query(`SELECT status FROM review WHERE body LIKE ('%the%')`)
	if err!=nil{
		logrus.Error("Can not execute query mysql")
	}
	if res.Next(){
		mes = "yes"
	}
	res.Close()
	db.Close()
	sqlEnd := time.Now()
	elapse := sqlEnd.Sub(sqlStart)
	return mes,elapse
}

func main() {

	mes1,time1 := es()
	mes2,time2 := mysql()
	fmt.Println("Time use ES: ",time1, " result: ",mes1)
	fmt.Println("Time use Mysql: ",time2, " result: ",mes2)

}