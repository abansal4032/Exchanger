package dal

import (
	"Exchanger/models"
	"Exchanger/server/dbclient"
	"database/sql"
	"errors"
	"github.com/nu7hatch/gouuid"
	"log"
)

const (
	GET_USERS = ` select user_id UserId, name Name, contact_number Contact, email Email, location Location, credits Credits, registration_token Registration_Token from user `
)

func GetUsers(userId string) ([]models.User, error) {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := GET_USERS
	var rows *sql.Rows
	if userId == "" {
		rows, err = tx.Query(query)
	} else {
		query = query + " where user_id = ? "
		rows, err = tx.Query(query, userId)
	}
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Contact, &user.Email,
			&user.Location, &user.Credits, &user.RegistrationToken); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(users) == 0 {
		return nil, errors.New("not found")
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	id, _ := uuid.NewV4()
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("INSERT INTO `user` (`user_id`, `name`, `contact_number`, `email`, `location`, `credits`)"+
		"VALUES (?, ?, ?, ?, ?, ?)", id.String(), user.Name, user.Contact, user.Email, user.Location, user.Credits)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func UpdateUser(token string, userId string) error {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("update user set registration_token = ? where user_id = ? ", token, userId)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
