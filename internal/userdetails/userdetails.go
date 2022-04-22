package userdetails

import (
	"log"

	"github.com/sarishap/go-authentication/database"
	"github.com/sarishap/go-authentication/internal/users"
)

type UserDetails struct {
	ID      string
	Name    string
	Address string
	Phone   string
	User    *users.User
}

func (userdetail UserDetails) Save() int64 {
	var lastInsertId int64
	err := database.Db.QueryRow("INSERT INTO usersdetails (name,address,phone,user_id) VALUES($1,$2,$3,$4) RETURNING id", userdetail.Name, userdetail.Address, userdetail.Phone, userdetail.User.ID).Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row inserted!")
	return lastInsertId

}
func FetchData() []UserDetails {
	stmt, err := database.Db.Query("select UD.id, UD.name, UD.address,UD.phone,UD.user_id, U.username from usersdetails UD inner join users U on UD.user_id = U.id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var userdetails []UserDetails
	var username string
	var id string
	for stmt.Next() {
		var userDetail UserDetails
		err := stmt.Scan(&userDetail.ID, &userDetail.Name, &userDetail.Address, &userDetail.Phone, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		userDetail.User = &users.User{
			ID:       id,
			Username: username,
		}
		userdetails = append(userdetails, userDetail)
	}
	if err = stmt.Err(); err != nil {
		log.Fatal(err)
	}
	return userdetails

}
