package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
mysql> CREATE TABLE userinfo (
    -> uid INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    -> username VARCHAR(64) NULL DEFAULT NULL,
    -> department VARCHAR(64) NULL DEFAULT NULL,
    -> created DATE NULL DEFAULT NULL
    -> )Engine=Innodb DEFAULT CHARSET=utf8;
*/

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(192.168.1.115:3306)/test?charset=utf8")
	checkErr(err)

	// insert
	fmt.Println("###### insert ######")
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?, department=?, created=?")
	checkErr(err)

	res, err := stmt.Exec("majun", "technique", "2018-06-05")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	fmt.Println("###### update ######")
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")

	res, err = stmt.Exec("songyue", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	fmt.Println("###### query ######")
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	fmt.Println("###### delete ######")

	stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
