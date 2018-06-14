package dal

import (
	"log"
	"Exchanger/models"
	"Exchanger/server/dbclient"
	"database/sql"
)
//
//type Requests struct {
//	RequestID        string `json:"requestId"`
//	EntityID         string `json:"entityId"`
//	Requester        string `json:"requester"`
//	Intent           string `json:"intent"`
//	DurationInDays   int    `json:"durationInDays"`
//	Status           string `json:"status"`
//	RequesterComment string `json:"requesterComment"`
//	OwnerComment     string `json:"ownerComment"`
//}

const (
	GET_REQUESTS = ` select request_id RequestId, entity_id EntityId, requester Requester, intent Intent, duration_in_days DurationInDays, status Status, requester_comment RequesterComment, owner_comment OwnerComment from requests `
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
	return requests, nil
}
