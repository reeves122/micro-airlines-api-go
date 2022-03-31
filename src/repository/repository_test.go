package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/reeves122/micro-airlines-api-go/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testPlayer = model.Player{
	Username: "testPlayer",
	Balance:  int64(1000),
}

func setUp() *Repository {
	id, _ := uuid.NewRandom()
	repo, err := NewRepository(fmt.Sprintf("%s.db", id))
	if err != nil {
		panic(err)
	}
	return repo
}

func tearDown(repo *Repository) {
	if err := repo.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(repo.FileName); err != nil {
		panic(err)
	}
}

func Test_NewRepository(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	assert.NotNil(t, repo)
}

func Test_InitializeDatabase(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	assert.NoError(t, initializeDatabase(repo))
}

func Test_Close(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	assert.NoError(t, repo.Close())
}

func Test_HealthCheck(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	assert.NoError(t, repo.HealthCheck())
}

func Test_AddPlayer(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	assert.NoError(t, repo.AddPlayer(testPlayer))
}

func Test_GetAllPlayers(t *testing.T) {
	repo := setUp()
	defer tearDown(repo)
	_ = repo.AddPlayer(testPlayer)
	results := repo.GetAllPlayers()
	assert.Equal(t, testPlayer.Username, results[0].Username)
	assert.Equal(t, testPlayer.Balance, results[0].Balance)
}
