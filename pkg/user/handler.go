package user

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/internal/database"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"go.uber.org/zap"
)

// insertUser inserts the user in the database
func insertUser(user *User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING user_id`

	if err := db.QueryRow(
		sqlStatement,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID); err != nil {
		log.L.Fatal("unable to execute the query", zap.Error(err))
	}

	log.L.Info("created user record", zap.Int64("userID", user.ID))

	return nil
}

// getUser gets only one user from the DB by its user_id
func getUserByID(ID int64) (*User, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	sqlStatement := `SELECT * FROM users WHERE user_id=$1`

	user := New()

	// execute the sql statement
	row := db.QueryRow(sqlStatement, ID)

	// unmarshal the row object to user
	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
		return nil, fmt.Errorf(
			"unable to execute `%s` with user_id %v: %v",
			sqlStatement, ID, err,
		)
	}

	return user, nil
}

// getAllUsers returns a list with the all users registered
func getAllUsers() ([]*User, error) {
	var users []*User

	db, err := database.Connect()
	if err != nil {
		return users, err
	}

	defer db.Close()

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return users, fmt.Errorf(
			"unable to execute `%s`: %v",
			sqlStatement, err,
		)
	}

	defer rows.Close()

	for rows.Next() {
		user := New()

		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			return users, fmt.Errorf("unable to scan a user row: %v", err)
		}

		users = append(users, user)
	}

	return users, err
}

// updateUser changes the user row in the database
// according to the new given User model
func updateUser(user *User) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `UPDATE users SET username=$2, email=$3, password=$4 WHERE user_id=$1`

	res, err := db.Exec(sqlStatement, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf(
			"unable to execute `%s` with user_id `%v`: %v",
			sqlStatement, user.ID, err,
		)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed checking the affected rows: %v", err)
	}

	log.L.Info("user rows updated", zap.Int64("total", rowsAffected))

	return nil
}

// deleteUser deletes an user from DB
func deleteUser(ID int64) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlStatement := `DELETE FROM users WHERE user_id=$1`

	res, err := db.Exec(sqlStatement, ID)
	if err != nil {
		return fmt.Errorf(
			"unable to execute `%s` with user_id `%v`: %v",
			sqlStatement, ID, err,
		)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed checking the affected rows: %v", err)
	}

	log.L.Info("user rows deleted", zap.Int64("total", rowsAffected))

	return nil
}
