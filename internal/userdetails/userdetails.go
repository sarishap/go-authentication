package userdetails

import (
	"log"

	"github.com/sarishap/go-authentication/database"
	"github.com/sarishap/go-authentication/users"
)

type UserDetails struct {
	ID      string
	Name    string
	Address string
	Phone   string
	User    *users.User
}

func (userdetail UserDetails) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO userdetails (name, address, phone) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(userdetail.Name, userdetail.Address, userdetail.Phone)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id

}

func FetchData() []UserDetails {
	stmt, err := database.Db.Prepare("select * from userdetails")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userdetails []UserDetails
	for rows.Next() {
		var userdetail UserDetails
		err := rows.Scan(&userdetail.ID, &userdetail.Name, &userdetail.Address, &userdetail.Phone)
		if err != nil {
			log.Fatal(err)
		}
		userdetails = append(userdetails, userdetail)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return userdetails

}
