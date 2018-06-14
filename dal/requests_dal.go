package dal

import (
	"Exchanger/models"
	"Exchanger/server/dbclient"
	"database/sql"
	"github.com/nu7hatch/gouuid"
	"errors"
	"log"
	"strings"
)

const (
	GET_REQUESTS = ` select requests.request_id RequestId, requests.entity_id EntityId, requests.requester Requester, requests.intent Intent, requests.duration_in_days DurationInDays, requests.status Status, requests.requester_comment RequesterComment, requests.owner_comment OwnerComment from requests `
)

func GetRequests(requestId string) ([]models.Requests, error) {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := GET_REQUESTS
	var rows *sql.Rows
	if requestId == "" {
		rows, err = tx.Query(query)
	} else {
		query = query + " where request_id = ? "
		rows, err = tx.Query(query, requestId)
	}
	if err != nil {
		log.Fatal(err)
	}
	var requests []models.Requests
	defer rows.Close()
	for rows.Next() {
		var request models.Requests
		if err := rows.Scan(&request.RequestID, &request.EntityID, &request.Requester, &request.Intent, &request.DurationInDays,
			&request.Status, &request.RequesterComment, &request.OwnerComment); err != nil {
			log.Fatal(err)
		}
		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(requests) == 0 {
		return nil, errors.New("not found")
	}
	return requests, nil
}

func GetRequestsByOwner(ownerName string) ([]models.Requests, error) {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := GET_REQUESTS
	var rows *sql.Rows
	query = query + " join entity using(entity_id) join user on user.user_id=entity.owner where user.name = ? "
	rows, err = tx.Query(query, ownerName)
	if err != nil {
		log.Fatal(err)
	}
	var requests []models.Requests
	defer rows.Close()
	for rows.Next() {
		var request models.Requests
		if err := rows.Scan(&request.RequestID, &request.EntityID, &request.Requester, &request.Intent, &request.DurationInDays,
			&request.Status, &request.RequesterComment, &request.OwnerComment); err != nil {
			log.Fatal(err)
		}
		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(requests) == 0 {
		return nil, errors.New("not found")
	}
	return requests, nil
}

func GetRequestsByRequester(requesterName string) ([]models.Requests, error) {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := GET_REQUESTS
	var rows *sql.Rows
	query = query + " where requester = ? "
	rows, err = tx.Query(query, requesterName)
	if err != nil {
		log.Fatal(err)
	}
	var requests []models.Requests
	defer rows.Close()
	for rows.Next() {
		var request models.Requests
		if err := rows.Scan(&request.RequestID, &request.EntityID, &request.Requester, &request.Intent, &request.DurationInDays,
			&request.Status, &request.RequesterComment, &request.OwnerComment); err != nil {
			log.Fatal(err)
		}
		requests = append(requests, request)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(requests) == 0 {
		return nil, errors.New("not found")
	}
	return requests, nil
}

func CreateRequest(request *models.Requests) error {
	id, _ := uuid.NewV4()
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec("INSERT INTO `requests` (`request_id`,`entity_id`,`requester`,`intent`,`duration_in_days`,`status`,`requester_comment`) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?)", id.String(), request.EntityID, request.Requester, request.Intent, request.DurationInDays, request.Status, request.RequesterComment)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func UpdateRequest(req *models.Requests, requstId string) error {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	status := req.Status
	_, err = tx.Exec("update requests set status = ?, owner_comment = ? where request_id = ? ", status, req.OwnerComment, requstId)
	if err != nil {
		return err
	}
	if strings.ToLower(status) == "approved" {
		request, err := GetRequests(requstId)
		if err != nil {
			return err
		}
		err = UpdateBorrower(request[0].EntityID, request[0].Requester, "Alloted")
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
