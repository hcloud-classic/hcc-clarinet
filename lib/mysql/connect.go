package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Needed for connect mysql
	"hcc/clarinet/lib/logger"
	"strconv"
)

// Db : Pointer of mysql connection
var Db *sql.DB

// Init : Initialize mysql connection
func Init(user string, password string, address string, port int64) error {
	var err error

	Db, err = sql.Open("mysql",
		user+":"+password+"@tcp("+
			address+":"+strconv.Itoa(int(port))+")/"+
			"piccolo"+"?parseTime=true")
	if err != nil {
		logger.Logger.Println(err)
		return err
	}

	err = Db.Ping()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}

	logger.Logger.Println("Connected to MySQL database")

	return nil
}

// End : Close mysql connection
func End() {
	if Db != nil {
		_ = Db.Close()
	}
}
