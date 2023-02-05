package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chack1920/tdriver/v3"
	"gorm.io/gorm"
)

type Data struct {
	TS    time.Time
	Value float64
}

func main() {
	db := connect()
	showTable(db)
	Table(db)

}
func connect() *gorm.DB {
	dsn := "root:taosdata@/tcp(192.168.123.96:6030)/orm_test?loc=Local"
	db, err := gorm.Open(tdriver.Open(dsn))
	if err != nil {
		log.Fatalf("unexpected error:%v", err)
	}
	db = db.Debug()
	return db
}
func showTable(db *gorm.DB) {
	type Result struct {
		Table_Name   string
		Created_Time time.Time
		Columns      int
		Stable_Name  string
		Uid          int64
		Tid          int16
		VgId         int16
	}
	var stu []Result
	rows := db.Raw("SHOW TABLES").Scan(&stu)
	tbnames := make([]string, len(stu))
	for i := 0; i < len(stu); i++ {
		tbnames[i] = stu[i].Table_Name

	}
	fmt.Println(tbnames)
	if rows != nil {

	}

}
func Table(db *gorm.DB) {

	querySQL := fmt.Sprintf("DESCRIBE %s", "tb1")
	rows, err := db.Raw(querySQL).Rows()
	if err != nil {
	}
	defer rows.Close()
	types, err := rows.ColumnTypes()
	if err != nil {
	}
	for i, columnType := range types {
		fmt.Println(columnType)
		fmt.Println(i)
	}

}
