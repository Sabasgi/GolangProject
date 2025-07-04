package handlers

import (
	"encoding/json"
	"log"
	models "repogin/internal/models/masters"
	"testing"
)

func TestGetLabs(t *testing.T) {
	router := GetTestHandlerRouter()
	lab := models.Lab{
		LabID:   1,
		LabName: "Test Lab",
	}
	labBytes, err := json.Marshal(&lab)
	if err != nil {
		log.Println("ERROR : TestGetLabs", err)
	}
	w := router.PerformRouteTest("GET", "c/lab/getall", string(labBytes))
	log.Println("WWWWWWWWWWWWW - ", w)

}
