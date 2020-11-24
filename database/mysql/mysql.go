package mysql

import (
	"github.com/panglove/BaseServer/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MysqlDB *sql.DB

func Init(db *config.MySql) {
	fmt.Println("open mysql")

	if !db.Enable {
		return
	}
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", db.USER, db.PSWD, db.HOST, db.PORT, db.DB, "utf8")
	MysqlDBInstall, err := sql.Open("mysql", dbDSN)

	if err != nil {
		log.Fatal(err)
	}
	MysqlDB = MysqlDBInstall
	fmt.Println("mysql connect " + dbDSN)

}

func Select(sqlStr string, k ...interface{}) ([]map[string]interface{}, bool) {

	rows, err := MysqlDB.Query(sqlStr, k...)

	if err != nil {
		fmt.Println("get query rows err:", err)
		return nil, false
	}

	rowsStrings, err := rows.Columns()

	rowsLength := len(rowsStrings)

	if err != nil {
		fmt.Println("get rows err:", err)
		return nil, false
	}

	if err != nil {
		return nil, false
	}

	var list []map[string]interface{}

	for rows.Next() {

		values := make([]interface{}, rowsLength)

		columnPointers := make([]interface{}, rowsLength)
		for i := 0; i < rowsLength; i++ {
			values[i] = &columnPointers[i]
		}

		rows.Scan(values...)

		valueItem := make(map[string]interface{})

		for i, key := range rowsStrings {

			valueItem[key] = *values[i].(*interface{})

			switch val := valueItem[key].(type) {
			case byte:
				valueItem[key] = val
				break
			case []byte:
				v := string(val)
				switch v {
				case "\x00": // 处理数据类型为bit的情况
					valueItem[key] = 0
				case "\x01": // 处理数据类型为bit的情况
					valueItem[key] = 1
				default:
					valueItem[key] = v
					break
				}
				break
			default:
				valueItem[key] = val
			}

		}
		list = append(list, valueItem)

	}
	return list, true

}
