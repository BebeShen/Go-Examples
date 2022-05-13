package main

import (
	"fmt"
	"time"
	
    . "app/pkg"
	. "app/expo"
	. "app/db"
)

func main ()  {
	Connect()
	CreateTable(DB)
	// stmt, err := DB.Prepare("INSERT INTO notification(user_id, token, device) values(?,?,?)")
	// CheckErr(err)
	// res, err := stmt.Exec("8839", "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]", "android")
    // CheckErr(err)

    // id, err := res.LastInsertId()
    // CheckErr(err)

    // fmt.Println(id)


	// user := Find(DB, 8840)
	// fmt.Printf("user_id: %d, token: %s, settings: %s\n", user.User_id, user.Token, user.Settings)

	// user.User_id = user.User_id + 1
	// user.Token = "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]"

	// if status := Insert(DB, user); status == "success" {
	// 	fmt.Println("[+] Insert Success!")
	// }
	// user.Settings = "1,1,1,0,1,1,1"
	// user.Device = "iOS"
	// user.Updated = time.Now()
	// Update(DB, user.User_id, user)

	fmt.Printf("Test expo token %s\n",time.Now().Format("YYYY-MM-DD"))
	// userList := FindAll(DB)
	// for _, u := range userList {
	// 	fmt.Printf("user_id: %d, token: %s, device: %s, settings: %s, created: %s. updated: %s\n", u.User_id, u.Token, u.Device, u.Settings, u.Created, u.Updated)
	// }
	var tmList TokensWithMessage
	tmList.Token = append(tmList.Token, "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]")
	tmList.Message = "Expo Push Test for golang (struct, slice)"
	u1, err := FindOne(DB, uint32(8839))
	fmt.Println("8839", u1, err)
	if _, err := FindOne(DB, uint32(8841)) ; err == ErrorUserNotExist {
		// insertUser := new(User)
		insertUser := new(User)
		insertUser.Init(8841, "ExponentPushToken[6hv555HcNo7iNLhnPt4Y9a]")
		if status := Insert(DB, insertUser) ; status != "success" {
			panic("Fail on INSERT USER")
		}
	}


	userList := FindAll(DB)
	for _, u := range userList {
		fmt.Printf("user_id: %d, token: %s, device: %s, settings: %s, created: %s. updated: %s\n", u.User_id, u.Token, u.Device, u.Settings, u.Created, u.Updated)
	}
	SendPushNotificationToUser(tmList)
	DB.Close()
}