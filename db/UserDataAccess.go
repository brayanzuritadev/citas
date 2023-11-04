package db

import (
	"database/sql"
	"fmt"

	"github.com/brayanzuritadev/citas/models"
	"github.com/brayanzuritadev/citas/tools"
)

func GetUser(email string) (models.User, bool) {
	var user models.User

	query := "SELECT UserId, FirstName, LastName, DateBirth, Email, Password, Avatar FROM [User] WHERE Email = @Email"

	rows, err := SQLDB.Query(query, sql.Named("Email", email))

	if err != nil {
		fmt.Println("Error executing SQL query:", err)
		return user, false
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.DateBirth, &user.Email, &user.Password, &user.Avatar)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return user, false
		}
		fmt.Println("User found:", user)
		return user, true
	}

	fmt.Println("No user found with email:", email)
	return user, false
}

func InsertUser(u models.User) (int, bool, error) {
	fmt.Println(u.Password)
	u.Password, _ = tools.PasswordEncrypt(u.Password)
	fmt.Println(u.Password)
	var userId int

	procedureName := "sp_InsertUser"

	query := fmt.Sprintf("EXEC %s @FirstName, @LastName, @DateBirth, @Email, @Password, @Avatar", procedureName)

	stmt, err := SQLDB.Prepare(query)
	if err != nil {
		fmt.Println("aqui esta el error" + err.Error())
		return 0, false, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("The statement was not closed")
		}
	}(stmt)

	err = stmt.QueryRow(
		sql.Named("FirstName", u.FirstName),
		sql.Named("LastName", u.LastName),
		sql.Named("DateBirth", u.DateBirth),
		sql.Named("Email", u.Email),
		sql.Named("Password", u.Password),
		sql.Named("Avatar", u.Avatar)).Scan(&userId)

	if err != nil {
		fmt.Println("aqui esta el error2" + err.Error() + u.FirstName)
		return 0, false, err
	}
	return userId, true, nil
}
