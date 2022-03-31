package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/reeves122/micro-airlines-api-go/mocks/mock_repository"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func Test_HealthCheck(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_repository.NewMockIRepository(ctrl)
		mockRepo.EXPECT().HealthCheck().Return(nil)
		resp := httptest.NewRecorder()
		HealthCheck(mockRepo, resp, nil)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("failed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockRepo := mock_repository.NewMockIRepository(ctrl)
		mockRepo.EXPECT().HealthCheck().Return(fmt.Errorf("some error"))
		resp := httptest.NewRecorder()
		HealthCheck(mockRepo, resp, nil)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func Test_respondJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		resp := httptest.NewRecorder()
		respondJSON(resp, 200, nil)
	})

	t.Run("failed", func(t *testing.T) {
		var invalidPayload = make(chan int)
		resp := httptest.NewRecorder()
		respondJSON(resp, 0, invalidPayload)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}
