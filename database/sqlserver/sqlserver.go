package sqlserver

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/panglove/BaseServer/config"
	"log"
	"strings"
)

type MssqlDB struct {
	*sql.DB
	NowSql string
	Table  string
}

func New(db *config.MsSql) *MssqlDB {
	fmt.Println("open mssql")

	if !db.Enable {
		return nil
	}
	dbDSN := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", db.HOST, db.PORT, db.DB, db.USER, db.PSWD)
	MssqlDBInstall, err := sql.Open("mssql", dbDSN)

	if err != nil {
		log.Fatal(err)
	}
	mdb := new(MssqlDB)
	mdb.DB = MssqlDBInstall
	fmt.Println("mssql connected " +db.HOST)
	return mdb

}

func (mdb *MssqlDB) DTable(tableName string) *MssqlDB {
	mdb.NowSql = ""
	mdb.Table = tableName
	return mdb
}
func (mdb *MssqlDB) DSelect(what ...string) *MssqlDB {
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
func (mdb *MssqlDB) DDelete() *MssqlDB {

	selectSql := fmt.Sprintf("delete from %s", mdb.Table)

	mdb.NowSql = selectSql
	return mdb
}

//like []string{"id=1","name=where"}

func (mdb *MssqlDB) DUpdate(what ...string) *MssqlDB {
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

func (mdb *MssqlDB) DInsert(what []string, data ...string) *MssqlDB {
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
func (mdb *MssqlDB) DWhereAnd(what ...string) *MssqlDB {
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
func (mdb *MssqlDB) DWhereOr(what ...string) *MssqlDB {
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

func (mdb *MssqlDB) DOrderBy(what ...string) *MssqlDB {
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
func (mdb *MssqlDB) DLimit(what ...string) *MssqlDB {
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
func (mdb *MssqlDB) DStrEqual(i string, k string) string {

	return i + "='" + k + "'"
}
func (mdb *MssqlDB) DIntEqual(i string, k string) string {

	return i + "=" + k + ""
}

func (mdb *MssqlDB) DExec() ([]map[string]interface{}, bool) {

	if len(mdb.NowSql) > 0 {
		if strings.Index(mdb.NowSql, "select") == 0 {
			return mdb.Select(mdb.NowSql)
		} else {
			_, err := mdb.Exec(mdb.NowSql)
			return nil, err == nil
		}
	}
	return nil, false
}

func (mdb *MssqlDB) Select(sqlStr string, k ...interface{}) ([]map[string]interface{}, bool) {

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
