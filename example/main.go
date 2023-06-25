package main

import (
	"fmt"
	"github.com/chack1920/tdriver/v3/clause/create"
	"log"
	"time"

	"github.com/chack1920/tdriver/v3"
	"gorm.io/gorm"
)

type Data struct {
	TS    time.Time
	Value float64
}
type AQI struct {
	TS   time.Time
	PM25 float64
}

func main() {
	db := connect()
	//result := map[string]interface{}{}
	// var d AQI
	// err := db.Table("aqi_866262045704805").Find(&d).Error
	// if err != nil {
	// 	log.Fatalf("找到数据 %v", err)
	// } // 返回找到的记录数
	resultWindowMax := SelectQuery(db, "stb_aqis")
	fmt.Println(resultWindowMax)
	//fmt.Println(&d)
	showTable(db)
	Table(db)

}
func SelectQuery(db *gorm.DB, tableName string) []map[string]interface{} {
	var result []map[string]interface{}
	err := db.Table(tableName).Find(&result)
	fmt.Println(err.RowsAffected)
	if err.Error != nil {
		log.Fatalf("aggregate query error %v", err.Error)
	}
	return result
}
func createTableUsingStable(db *gorm.DB) {
	// create table using sTable
	table := create.NewTable("tb_1", true, nil, "stb_1", map[string]interface{}{
		"tbn": "tb_1",
	})
	err := db.Table("tb_1").Clauses(create.NewCreateTableClause([]*create.Table{table})).Create(map[string]interface{}{}).Error
	if err != nil {
		log.Fatalf("create table error %v", err)
	}
}
func connect() *gorm.DB {
	dsn := "root:taosdata@/tcp(192.168.123.96:6030)/aqis?loc=Local"
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

	querySQL := fmt.Sprintf("DESCRIBE %s", "aqi_866262045704805")
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
