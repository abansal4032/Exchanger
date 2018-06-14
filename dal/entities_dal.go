package dal

import (
	"Exchanger/server/dbclient"
	"Exchanger/models"
	"log"
	"database/sql"
	"errors"
)

func GetAllEntitites(entityID string) ([]models.Entity, error) {
	tx, err := dbclient.NewTransaction()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var rows *sql.Rows
	if entityID == "" {
		rows, err = tx.Query("SELECT entity_id, entity_name, entity_type, owner, action_type, status, price, borrower, location from entity")
	} else {
		rows, err = tx.Query("SELECT entity_id, entity_name, entity_type, owner, action_type, status, price, borrower, location from entity where entity_id = ?", entityID)
	}
	if err != nil {
		log.Fatal(err)
	}
	var entities []models.Entity
	defer rows.Close()
	for rows.Next() {
		var entity models.Entity
		entity.Attributes = make(map[string]string)
		if err := rows.Scan(&entity.EntityID, &entity.Name, &entity.Type, &entity.Owner, &entity.Action, &entity.Status, &entity.Price, &entity.Borrower, &entity.Location); err != nil {
			log.Fatal(err)
		}
		tx2, err := dbclient.NewTransaction()
		if err != nil {
			return nil, err
		}
		defer tx2.Rollback()
		attributes, err := tx2.Query("SELECT attribute_key, attribute_value from entity_attributes where deleted_at = 0 and entity_id = ?", entity.EntityID)
		if err != nil {
			log.Fatal(err)
		}
		defer attributes.Close()
		for attributes.Next() {
			var key string
			var val string
			if err := attributes.Scan(&key, &val); err != nil {
				log.Fatal(err)
			}
			entity.Attributes[key] = val
		}
		entities = append(entities, entity)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(entities) == 0 {
		return nil, errors.New("not found")
	}
	return entities, nil
}