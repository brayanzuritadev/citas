package db

import (
	"database/sql"
	"fmt"

	"github.com/brayanzuritadev/citas/models"
	"github.com/brayanzuritadev/citas/tools"
)

func GetUser(email string) (models.User, error) {
	var user models.User

	query := "SELECT * FROM [User] WHERE Email = @Email"

	rows, err := SQLDB.Query(query, sql.Named("Email", email))

	if err != nil {
		fmt.Println("Error executing SQL query:", err)
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.BirthDate, &user.Email, &user.Password, &user.Avatar, &user.IsDeleted, &user.IsLocked)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return user, err
		}
		fmt.Println("User found:", user)
		return user, err
	}

	fmt.Println("No user found with email:", email)
	return user, nil
}

func InsertUser(u models.User) (int, bool, error) {
	fmt.Println(u.Password)
	u.Password, _ = tools.PasswordEncrypt(u.Password)
	fmt.Println(u.Password)
	var userId int

	procedureName := "sp_InsertUser"

	query := fmt.Sprintf("EXEC %s @FirstName, @LastName, @BirthDate, @Email, @Password, @Avatar, @IsDeleted, @IsLocked", procedureName)

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
		sql.Named("BirthDate", u.BirthDate),
		sql.Named("Email", u.Email),
		sql.Named("Password", u.Password),
		sql.Named("Avatar", u.Avatar),
		sql.Named("IsDeleted", u.IsDeleted),
		sql.Named("IsLocked", u.IsLocked),
	).Scan(&userId)

	if err != nil {
		fmt.Println("aqui esta el error2" + err.Error() + u.FirstName)
		return 0, false, err
	}
	return userId, true, nil
}
