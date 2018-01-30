package main
import(
	"bytes"
	 _ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

const(
	N = 15
)

func openDB()(db *sql.DB){
	var err error
	db,err = sql.Open(sqls, user+password+database+charset)
	db.SetMaxIdleConns(N)
	db.SetMaxOpenConns(N)
	db.SetConnMaxLifetime(N)
	writeError(err)
	return
}


	var(
		user = getKey(bdd,"user")
		sqls = getKey(bdd,"driver")
		password = ":"+getKey(bdd,"pwd")+"@/"
		charset = getKey(bdd,"utf")
		database = getKey(bdd,"database")
	)
	


func concat(init string,ar ...string)(oi string){
		are := bytes.NewBufferString(init)
		for _,ok := range ar{
			are.Write([]byte(ok))
		}
		return are.String()
}

func dbCommand(a string){
	db := openDB()
	_,err := db.Exec(a)
	dbError(err)
}

func dbError(er error)(err error){
	if er != nil{
		log.Fatal(er)
	}
	return er
}

func dbPrepare(a string)(ilo *sql.Stmt,err error){
	db := openDB()
	stmt, err := db.Prepare(a)
	return  stmt, dbError(err)
}

func dbExec(stmt *sql.Stmt, ar ...interface{})(a sql.Result,err error){
	res, err := stmt.Exec(ar...)
	return res,dbError(err)
}

func dbQueryOne(stmt *sql.Stmt, ar ...interface{})(a *sql.Rows,err error){
	res, err := stmt.Query(ar...)
	return res,dbError(err)
}

func dbQueryRows(a string,ar ...interface{})(io int,erer error){
	var res int
	db := openDB()
	err := db.QueryRow(a).Scan(&res)
	return res,dbError(err)
}

func dbQuery(a string)(io *sql.Rows, err error){
	db := openDB()
	rows, err := db.Query(a)
	return rows,dbError(err)
}




func table(){
	db := openDB()
	_,err:=db.Exec("CREATE TABLE  IF NOT EXISTS  user (login varchar(100) UNIQUE NOT NULL,mail varchar(300) NOT  NULL UNIQUE,password varchar(300) NOT NULL)")
	writeError(err)
}
func init(){
	db,err := sql.Open(sqls, user+password+charset)
	writeError(err)
	_,err = db.Exec("CREATE DATABASE IF NOT EXISTS "+ database)
	writeError(err)
	table()
}
