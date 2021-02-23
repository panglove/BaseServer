package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/panglove/BaseServer/config"
	"log"
	"strings"
)

type MysqlDB struct {
	*sql.DB
	NowSql string
	Table  string
}

func New(db *config.MySql) *MysqlDB {
	fmt.Println("open mysql")

	if !db.Enable {
		return nil
	}
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", db.USER, db.PSWD, db.HOST, db.PORT, db.DB, "utf8")
	MysqlDBInstall, err := sql.Open("mysql", dbDSN)

	if err != nil {
		log.Fatal(err)
	}
	mdb := new(MysqlDB)
	mdb.DB = MysqlDBInstall
	fmt.Println("mysql connect " + dbDSN)
	return mdb

}

func (mdb *MysqlDB) DTable(tableName string) *MysqlDB {
	mdb.NowSql = ""
	mdb.Table = tableName
	return mdb
}
func (mdb *MysqlDB) DSelect(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + ","
		} else {
			sql += v

		}
	}
	selectSql := fmt.Sprintf("select %s from %s", sql, mdb.Table)

	mdb.NowSql = selectSql
	return mdb

}
func (mdb *MysqlDB) DDelete() *MysqlDB {

	selectSql := fmt.Sprintf("delete from %s", mdb.Table)

	mdb.NowSql = selectSql
	return mdb
}

//like []string{"id=1","name=where"}

func (mdb *MysqlDB) DUpdate(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + ","
		} else {
			sql += v

		}
	}
	selectSql := fmt.Sprintf("update %s set %s", mdb.Table, sql)

	mdb.NowSql = selectSql
	return mdb
}

//like []string{"id","name"},[]interface{1,"12"}

func (mdb *MysqlDB) DInsert(what []string, data ...string) *MysqlDB {
	sql := "("
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + ","
		} else {
			sql += v

		}
	}
	sql += ")"

	datasql := "("
	for i, v := range data {
		if i != len(data)-1 {
			datasql += v + ","
		} else {
			datasql += v

		}
	}
	datasql += ")"
	selectSql := fmt.Sprintf("insert into  %s %s values %s", mdb.Table, sql, datasql)
	mdb.NowSql = selectSql
	return mdb
}

//like []string{"id=1","name=where"}
func (mdb *MysqlDB) DWhereAnd(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + " and "
		} else {
			sql += v

		}
	}
	whereSql := fmt.Sprintf(mdb.NowSql+" where %s", sql)
	mdb.NowSql = whereSql
	return mdb
}

//like []string{"id=1","name=where"}
func (mdb *MysqlDB) DWhereOr(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + " or "
		} else {
			sql += v

		}
	}
	whereSql := fmt.Sprintf(mdb.NowSql+" where %s", sql)
	mdb.NowSql = whereSql
	return mdb
}

func (mdb *MysqlDB) DOrderBy(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + ","
		} else {
			sql += v

		}
	}
	whereSql := fmt.Sprintf(mdb.NowSql+" order by %s", sql)
	mdb.NowSql = whereSql
	return mdb
}
func (mdb *MysqlDB) DLimit(what ...string) *MysqlDB {
	sql := ""
	for i, v := range what {
		if i != len(what)-1 {
			sql += v + ","
		} else {
			sql += v

		}
	}
	whereSql := fmt.Sprintf(mdb.NowSql+" limit %s", sql)
	mdb.NowSql = whereSql
	return mdb
}
func (mdb *MysqlDB) DStrEqual(i string,k string) string {

	return  i+"='"+k+"'"
}
func (mdb *MysqlDB) DIntEqual(i string,k string) string {

	return  i+"="+k+""
}

func (mdb *MysqlDB) DExec() ([]map[string]interface{}, bool) {

	if len(mdb.NowSql) > 0 {
		if strings.Index(mdb.NowSql, "select") == 0 {
			return mdb.Select(mdb.NowSql)
		} else {
			_, err := mdb.Exec(mdb.NowSql)
			log.Println(err)
			return nil, err == nil
		}
	}
	return nil, false
}

func (mdb *MysqlDB) Select(sqlStr string, k ...interface{}) ([]map[string]interface{}, bool) {

	rows, err := mdb.Query(sqlStr, k...)

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
