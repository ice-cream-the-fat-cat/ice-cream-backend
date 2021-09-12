package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gardens_models "github.com/ice-cream-backend/models/v1/gardens"
	gardens_router "github.com/ice-cream-backend/routes/v1/gardens"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateGardens(t *testing.T)  {
	utils.LoadEnv(true)

	payload := []byte(`
	{"name":"Test Garden","description":"Test Description","fireBaseUserId":"testUUID","gardenCategoryId":"613dd5bf3147c8eac03bcec1"}`)

	req, err := http.NewRequest("POST", "/api/v1/gardens", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gardens_router.CreateGardens)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var createdGarden gardens_models.Gardens
	json.Unmarshal(rr.Body.Bytes(), &createdGarden)

	if createdGarden.Name != "Test Garden" {
		t.Errorf("Expected garden's name to be 'Test Garden'. Got '%v'", createdGarden.Name)
	}

	if createdGarden.FireBaseUserId != "testUUID" {
		t.Errorf("Expected garden's fireBaseUserId to be 'testUUID'. Got '%v'", createdGarden.FireBaseUserId)
	}

	if createdGarden.Description != "Test Description" {
		t.Errorf("Expected garden's description to be 'Test Description'. Got '%v'", createdGarden.Description)
	}

	if createdGarden.GardenCategoryId.String() != "ObjectID(\"613dd5bf3147c8eac03bcec1\")" {
		t.Errorf("Expected garden's gardenCategoryId to be '613dd5bf3147c8eac03bcec1'. Got '%v'", createdGarden.GardenCategoryId)
	}

	if createdGarden.GardenCategory.Name == "" || createdGarden.GardenCategory.ID == primitive.NilObjectID {
		t.Errorf("Expected garden's gardenCategory to be populated. Got '%v'", createdGarden.GardenCategory)
	}
}