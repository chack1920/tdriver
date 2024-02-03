package main

import (
	"fmt"
	"github.com/chack1920/tdriver/v3"
	"github.com/chack1920/tdriver/v3/clause/create"
	"github.com/chack1920/tdriver/v3/clause/using"
	"gorm.io/gorm"
	"log"
	"time"
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
	val := map[string]interface{}{
		"snub":       "866262045704805",            //设备名称
		"typesn":     "MY-2022",                    //设备型号
		"projectid":  "xm32423423434",              //项目编号
		"provinceid": "430000",                     //省编号
		"cityid":     "430400",                     //市编号
		"areaid":     "430408",                     //区编号
		"supplierid": "gy23423dfdsf34234234234234", //供应商id
	}
	valel := map[string]interface{}{
		"ts":            time.Now(),
		"pm25":          101.5,
		"pm10":          100.2,
		"tsp":           43.2,
		"noise":         55.2,
		"temperature":   23,
		"humidity":      66,
		"wind_speed":    3,
		"winddirection": "东南",
		"wind_power":    3,
		"air_pressure":  1000.1, //气压
		"longitude":     26.907082,
		"latitude":      112.557941,
		"warning_state": 0,
	}
	aqiAutoInData(db, "stb_aqis", "866262045704805", val, valel)
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

/*func testmap(tableName string, typesn string, projectid string, provinceid string, cityid string, areaid string, supplierid string) {
	val := map[string]interface{}{
		"snub":       tableName,  //设备名称
		"typesn":     typesn,     //设备型号
		"projectid":  projectid,  //项目编号
		"provinceid": provinceid, //省编号
		"cityid":     cityid,     //市编号
		"areaid":     areaid,     //区编号
		"supplierid": supplierid, //供应商id
	}

}*/

// INSERT INTO tb_2 USING stb_1('tbn') TAGS('tb_2') (ts,value) VALUES ('2021-08-11 09:43:01.041',0.940509)
// automaticTableCreationWhenInsertingData(db, "tb_2", t1, randValue2)
func autoTableCreatInsertData(db *gorm.DB, tableName string, ts time.Time, value interface{}) {
	//automatic table creation when inserting data
	err := db.Table(tableName).Clauses(using.SetUsing("stb_1", map[string]interface{}{
		"tbn": tableName,
	})).Create(map[string]interface{}{
		"ts":    ts,
		"value": value,
	}).Error
	if err != nil {
		log.Fatalf("create table when insert data error %v", err)
	}
}

func aqiAutoInData(db *gorm.DB, sTable string, tableName string, val map[string]interface{}, value map[string]interface{}) {
	err := db.Table("aqi_" + tableName).Clauses(using.SetUsing(sTable, val)).Create(value).Error
	if err != nil {
		log.Fatalf("创建数据失败 %v", err)
	}
	//通过stable插入语句
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
