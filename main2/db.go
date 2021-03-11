import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
	)
	
	func main() {
	
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test-db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.ping()
	if err != nil {
	   log.Print(err)
	}