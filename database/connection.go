package database
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var DB *sql.DB
func init(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db1")
	if err!=nil{
		logrus.Error("Can not connect mysql")
	}


	DB = db
	DB.SetMaxOpenConns(50)
	DB.SetMaxIdleConns(50)
	DB.SetConnMaxLifetime(time.Second*1)
	DB.SetConnMaxIdleTime(time.Second*1)

}
func GetConnect() *sql.DB{
	return DB
}
