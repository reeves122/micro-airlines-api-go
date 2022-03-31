package handler

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/reeves122/micro-airlines-api-go/mocks/mock_repository"
	"github.com/reeves122/micro-airlines-api-go/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testPlayer = model.Player{
	Username: "testPlayer",
	Balance:  int64(1000),
}

func Test_GetAllPlayers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repository.NewMockIRepository(ctrl)
	mockRepo.EXPECT().GetAllPlayers().Return([]model.Player{
		{
			Username: "user1",
		},
		{
			Username: "user2",
		},
	})
	resp := httptest.NewRecorder()
	GetAllPlayers(mockRepo, resp, nil)
	assert.Equal(t, http.StatusOK, resp.Code)
	var returnedPlayers []model.Player
	if err := unmarshalResponse(resp, &returnedPlayers); err != nil {
		t.Error(err)
	}
	assert.Len(t, returnedPlayers, 2)
	assert.Equal(t, "user1", returnedPlayers[0].Username)
	assert.Equal(t, "user2", returnedPlayers[1].Username)
}

func Test_CreatePlayer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock_repository.NewMockIRepository(ctrl)

	t.Run("success", func(t *testing.T) {
		resp := httptest.NewRecorder()
		mockRepo.EXPECT().AddPlayer(testPlayer).Return(nil)
		req, _ := http.NewRequest("POST", "", structToReader(testPlayer))
		CreatePlayer(mockRepo, resp, req)
		assert.Equal(t, http.StatusCreated, resp.Code)
	})

	t.Run("bad body", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var body bytes.Buffer
		body.Write([]byte("foo"))
		req, _ := http.NewRequest("POST", "", &body)
		CreatePlayer(mockRepo, resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("database error", func(t *testing.T) {
		resp := httptest.NewRecorder()
		mockRepo.EXPECT().AddPlayer(testPlayer).Return(fmt.Errorf("some error"))
		req, _ := http.NewRequest("POST", "", structToReader(testPlayer))
		CreatePlayer(mockRepo, resp, req)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
