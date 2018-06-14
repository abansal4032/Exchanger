package dal

import (
	"log"
	"Exchanger/models"
	"Exchanger/server/dbclient"
	"database/sql"
)

const (
	GET_USERS = ` select user_id UserId, name Name, contact_number Contact, email Email, location Location, credits Credits from user `
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
			&user.Location, &user.Credits); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users, nil
}
