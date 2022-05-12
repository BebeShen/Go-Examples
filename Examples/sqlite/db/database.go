package db

import (
	"database/sql"
    "fmt"
    "time"

    . "db/pkg"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./notification.db")
	CheckErr(err)
    // global db connection object
    DB = db

	return db, err
}

func CreateTable(db *sql.DB) {
	// create table
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
}

func FindAll(db *sql.DB) []User {
    rows, err := db.Query("SELECT * FROM notification")
    // defer rows.Close()
    var u User
    var userList []User
    for rows.Next() {
        err = rows.Scan(&u.User_id, &u.Token, &u.Device, &u.Settings, &u.Created, &u.Updated)
        CheckErr(err)
        userList = append(userList, u)
        // fmt.Printf("user_id: %d, token: %s, settings: %s\n", u.User_id, u.Token, u.Settings)
    }
    return userList
}

func FindOne(db *sql.DB, user_id int) (user User) {
    result := db.QueryRow("SELECT * FROM notification WHERE user_id=?", user_id)
    // defer result.Close()
    result.Scan(&user.User_id, &user.Token, &user.Device, &user.Settings, &user.Created, &user.Updated)
    return user
}

func Insert(db *sql.DB, user User) string {
    // insert
    stmt, err := db.Prepare("INSERT INTO notification(user_id, token, device, settings, created, updated) values(?,?,?,?,datetime('now'),datetime('now'))")
    CheckErr(err)
    res, err := stmt.Exec(user.User_id, user.Token, user.Device, user.Settings, time.Now(), time.Now())
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func Update(db *sql.DB, userId int, user User) bool {
    // update
    stmt, err := db.Prepare("update notification set (token, settings, device, updated)=(?,?,?,?) where user_id=?")
    CheckErr(err)
    user.Updated = time.Now()
    res, err := stmt.Exec(user.Token, user.Settings, user.Device, user.Updated, user.User_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}

/*
func main() {
	// db, err := sql.Open("sqlite3", "./notification.db")
    db, err := connect()
    checkErr(err)
    fmt.Printf("type of db: %T\n", db)
    checkErr(err)
    createTable(db)

    

   

    // query
    findAll(db)

    user := find(db, 8839)
    fmt.Printf("user_id: %d, token: %s, settings: %s\n", user.User_id, user.Token, user.Settings)


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
*/

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}