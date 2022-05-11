package main

import (
	"database/sql"
    "fmt"
    "time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./notification.db")
    checkErr(err)
    fmt.Printf("type of db: %T\n", db)
    checkErr(err)
    //建立表
    sql_table := `
    CREATE TABLE IF NOT EXISTS notification(
        user_id INTEGER PRIMARY KEY,
        token VARCHAR(255) NULL,
        device VARCHAR(15) NULL,
        settings VARCHAR(15) DEFAULT '1,1,1,1,1,1,1',
        created DATE DEFAULT CURRENT_TIMESTAMP,
        updated DATE DEFAULT CURRENT_TIMESTAMP
    );
	`
    db.Exec(sql_table)

    // // insert
    // stmt, err := db.Prepare("INSERT INTO notification(user_id, token, device) values(?,?,?)")
    // checkErr(err)

    // res, err := stmt.Exec("8839", "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]", "android")
    // checkErr(err)

    // id, err := res.LastInsertId()
    // checkErr(err)

    // fmt.Println(id)

    // update
    stmt, err := db.Prepare("update notification set settings=? where user_id=?")
    checkErr(err)

    res, err := stmt.Exec("1,1,1,0,1,1,1", 8839)
    checkErr(err)
    
    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // query
    rows, err := db.Query("SELECT * FROM notification")
    checkErr(err)
    var user_id int
    var token string
    var device string
    var settings string
    var created time.Time
    var updated time.Time

    for rows.Next() {
        err = rows.Scan(&user_id, &token, &device, &settings, &created, &updated)
        checkErr(err)
        fmt.Printf("user_id: %d, token: %s, settings: %s\n", user_id, token, settings)
    }

    rows.Close()

    // // delete
    // stmt, err = db.Prepare("delete from userinfo where uid=?")
    // checkErr(err)

    // res, err = stmt.Exec(id)
    // checkErr(err)

    // affect, err = res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}