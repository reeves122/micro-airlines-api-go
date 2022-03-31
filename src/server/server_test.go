package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/reeves122/micro-airlines-api-go/model"
	"github.com/reeves122/micro-airlines-api-go/repository"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testPlayer = model.Player{
	Username: "testPlayer",
	Balance:  int64(1000),
}

func setUp() (*repository.Repository, *server) {
	id, _ := uuid.NewRandom()
	repo, err := repository.NewRepository(fmt.Sprintf("%s.db", id))
	if err != nil {
		panic(err)
	}
	svr := NewServer(repo)
	return repo, svr
}

func tearDown(repo *repository.Repository) {
	if err := repo.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(repo.FileName); err != nil {
		panic(err)
	}
}

func unmarshalResponse(resp *httptest.ResponseRecorder, destination interface{}) error {
	data, err := ioutil.ReadAll(resp.Result().Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &destination); err != nil {
		return fmt.Errorf("returned data is invalid JSON. Got: %s", data)
	}
	return nil
}

func structToReader(s interface{}) *bytes.Buffer {
	var body bytes.Buffer
	_ = json.NewEncoder(&body).Encode(s)
	return &body
}

func Test_GetHealthCheck(t *testing.T) {
	repo, svr := setUp()
	defer tearDown(repo)
	req, _ := http.NewRequest("GET", "/api/healthcheck", nil)
	resp := httptest.NewRecorder()

	svr.setRouters()
	svr.router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func Test_GetAllPlayers(t *testing.T) {
	repo, svr := setUp()
	defer tearDown(repo)

	_ = repo.AddPlayer(testPlayer)
	req, _ := http.NewRequest("GET", "/api/players", nil)
	resp := httptest.NewRecorder()

	svr.setRouters()
	svr.router.ServeHTTP(resp, req)

	var returnedPlayers []model.Player
	if err := unmarshalResponse(resp, &returnedPlayers); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, testPlayer.Username, returnedPlayers[0].Username)
	assert.Equal(t, testPlayer.Balance, returnedPlayers[0].Balance)
}

func Test_CreatePlayer(t *testing.T) {
	repo, svr := setUp()
	defer tearDown(repo)

	req, _ := http.NewRequest("POST", "/api/players", structToReader(testPlayer))
	resp := httptest.NewRecorder()
	svr.setRouters()
	svr.router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	results := repo.GetAllPlayers()
	assert.Equal(t, testPlayer.Username, results[0].Username)
	assert.Equal(t, testPlayer.Balance, results[0].Balance)
}
