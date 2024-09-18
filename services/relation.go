package services

import (
	"encoding/json"

	models "groupietracker/models"
	utils "groupietracker/utils"
)

func fetchRelationById(id string) ([]byte, error) {
	return utils.FetchGroupieTracker("relation/" + id)
}

func fetchRelations() ([]byte, error) {
	return utils.FetchGroupieTracker("relation")
}

func GetRelations() ([]models.Relation, error) {
	var relation []models.Relation
	relationsData, err := fetchRelations()
	if err != nil {
		return relation, err
	}
	json.Unmarshal(relationsData, &relation)
	return relation, nil
}

func GetRelationById(id string) (models.Relation, error) {
	var relation models.Relation
	relationData, err := fetchRelationById(id)
	if err != nil {
		return relation, err
	}
	json.Unmarshal(relationData, &relation)
	return relation, nil
}
