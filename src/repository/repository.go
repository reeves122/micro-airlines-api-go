package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/reeves122/micro-airlines-api-go/model"
)

type IRepository interface {
	AddPlayer(player model.Player) error
	GetAllPlayers() []model.Player
	HealthCheck() error
}

type Repository struct {
	db       *gorm.DB
	FileName string
}

// NewRepository initializes a new repository
func NewRepository(fileName string) (*Repository, error) {
	db, err := gorm.Open("sqlite3", fileName)
	if err != nil {
		return nil, err
	}

	repo := &Repository{
		db,
		fileName,
	}

	err = initializeDatabase(repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func initializeDatabase(repo *Repository) error {
	repo.db.AutoMigrate(
		&model.City{},
		&model.Job{},
		&model.Plane{},
		&model.Player{},
	)
	return nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) HealthCheck() error {
	err := r.db.DB().Ping()
	return err
}

func (r *Repository) AddPlayer(player model.Player) error {
	result := r.db.Model(&model.Player{}).Save(&player)
	return result.Error
}

func (r *Repository) GetAllPlayers() []model.Player {
	var players []model.Player
	r.db.Model(&model.Player{}).Find(&players)
	return players
}
