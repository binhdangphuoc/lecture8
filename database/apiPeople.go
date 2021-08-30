package database

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"lecture8/models"
	_ "lecture8/models"
	"net/http"
)
func Ping(){
	e := DB.Ping()
	if e!=nil {
		fmt.Println("loi connect mysql api")
	}
	fmt.Println("ok",DB)
	res, _ := DB.Query("SHOW TABLES")
	defer res.Close()
	var table string
	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
}

func GetPeopleId(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request get a People")
	para:=mux.Vars(r)
	rs,err := DB.Query(`SELECT id,name,birth_day,address,phone,sex,job FROM people WHERE id=?`,para["id"])
	defer rs.Close()
	if err!=nil{
		logrus.Error("Error query database GetBuyId")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	var people models.People
	if rs.Next(){
		var(
			id int
			name string
			birthDay string
			address string
			phone string
			sex string
			job string
		)
		err = rs.Scan(&id,&name,&birthDay,&address,&phone,&sex,&job)
		if err!=nil{
			logrus.Error("Error scan database GetPeopleId")
			_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
			return
		}
		people = models.People{Id: id,Name: name,BirthDay: birthDay,Address: address,Phone: phone,Sex: sex,Job: job}
	}else{
		logrus.Error("No found id people")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	logrus.Info("Success GetBuyId")
	_ = json.NewEncoder(w).Encode(&people)
}

func GetAllPeople(w http.ResponseWriter, r *http.Request){
	logrus.Info("Request get all People")
	Ping()
	var people []models.People
	rs,err := DB.Query(`SELECT id,name,birth_day,address,phone,sex,job FROM people`)
	defer rs.Close()
	if err!=nil{
		logrus.Error("Error query database get all people")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	for rs.Next(){
		var(
			id int
			name string
			birthDay string
			address string
			phone string
			sex string
			job string
		)
		err = rs.Scan(&id,&name,&birthDay,&address,&phone,&sex,&job)
		if err!=nil{
			logrus.Error("Error scan database get all people")
			_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
			return
		}
		people = append(people,models.People{Id: id,Name: name,BirthDay: birthDay,Address: address,Phone: phone,Sex: sex,Job: job})
	}
	logrus.Info("Success get all people")
	_ = json.NewEncoder(w).Encode(&people)
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request create People")
	var newPeople models.People
	err:= json.NewDecoder(r.Body).Decode(&newPeople)
	if err!=nil{
		logrus.Error("Error decode new people")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	fmt.Println(newPeople)
	rs, err := DB.Query(`INSERT INTO people (name,birth_day,address,phone,sex,job) 
		VALUES (?,?,?,?,?,?)`,newPeople.Name,newPeople.BirthDay,newPeople.Address,newPeople.Phone,newPeople.Sex,newPeople.Job)
	defer rs.Close()
	if err != nil {
		logrus.Error("Error query create new people")
		return
	}
	logrus.Info("Success create new people")
	_ = json.NewEncoder(w).Encode(http.StatusOK)
}

func UpdatePeople(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request update people")
	para := mux.Vars(r)
	rs, err := DB.Query(`SELECT id,name,birth_day,address,phone,sex,job FROM people WHERE id=?`,para["id"])
	defer rs.Close()
	if err != nil {
		logrus.Error("Error query get id to update people")
		return
	}
	var people models.People
	if rs.Next() {
		err = json.NewDecoder(r.Body).Decode(&people)
		if err != nil||people.Name=="" {
			logrus.Error("Request body has mistake")
		}
		rs2, err := DB.Query(`UPDATE people SET name=?,birth_day=?,address=?,phone=?,sex=?,job=? 
		WHERE id = ?`,people.Name,people.BirthDay,people.Address,people.Phone,people.Sex,people.Job,para["id"])
		defer rs2.Close()
		if err != nil {
			logrus.Error("Error query update people")
			return
		}
	} else {
		logrus.Info("No found people id from request")
		_ = json.NewEncoder(w).Encode(http.StatusBadRequest)
		return
	}
	logrus.Info("Success update people")
	_ = json.NewEncoder(w).Encode(http.StatusOK)
}

func DeletePeople(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Request delete people")
	para := mux.Vars(r)
	rs, err := DB.Query(`DELETE FROM people WHERE id=?`,para["id"])
	defer rs.Close()
	if err != nil {
		logrus.Error("Error query delete people")
		return
	}
	test,err := rs.Columns()
	if err != nil {
		logrus.Error("Error query delete people 2222222")
		return
	}
	logrus.Info(test)
	logrus.Info("Success delete people")
	_ = json.NewEncoder(w).Encode(http.StatusOK)

}


//INSERT INTO people (name,birth_day,address,phone,sex,job) VALUES ("Nguyễn Văn Bách",?,?,?,?,?)