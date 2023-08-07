package gettable

import (
	"github.com/golangast/goservershell/internal/dbsql/dbconn"
)

func Gettabledata(p string) []User {
	data, err := dbconn.DbConnection() //create db instance
	dbconn.ErrorCheck(err)

	//variables used to store data from the query
	var (
		id    string
		name  string
		email string
		users []User
		u     User
	)

	//get from database
	rows, err := data.Query("SELECT * FROM users WHERE name = ?", p)
	dbconn.ErrorCheck(err)
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		dbconn.ErrorCheck(err)

		//store into memory
		u = User{Id: id, Name: name, Email: email}
		users = append(users, u)
	}
	//close everything
	rows.Close()
	data.Close()
	return users

}

type User struct {
	Id    string
	Name  string
	Email string
}
