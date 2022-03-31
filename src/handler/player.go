package handler

import (
	"encoding/json"
	"github.com/reeves122/micro-airlines-api-go/model"
	"github.com/reeves122/micro-airlines-api-go/repository"
	"net/http"
)

func GetAllPlayers(repo repository.IRepository, w http.ResponseWriter, _ *http.Request) {
	players := repo.GetAllPlayers()
	respondJSON(w, http.StatusOK, players)
	return
}

func CreatePlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {
	player := model.Player{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&player); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := repo.AddPlayer(player); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, player)
}

//func GetPlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {}
//
//func UpdatePlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {}
//
//func DeletePlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {}
//
//func DisablePlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {}
//
//func EnablePlayer(repo repository.IRepository, w http.ResponseWriter, r *http.Request) {}
