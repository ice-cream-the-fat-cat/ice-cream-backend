package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	gardens_router "github.com/ice-cream-backend/routes/v1/gardens"
	"github.com/ice-cream-backend/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateGardens(t *testing.T)  {
	utils.LoadEnv()
	oid := primitive.NewObjectID().String()
	var jsonStr = []byte(`
		{"_id":"ObjectID("` + oid + `")","name":"Test Garden","description":"Test Description"}
		`)

	req, err := http.NewRequest("POST", "/api/v1/gardens", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(gardens_router.CreateGardens)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `
		{"_id":"` + oid + `","name":"Test Garden","description":"Test Description"}
	`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}