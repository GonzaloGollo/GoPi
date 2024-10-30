package user

import (
	"context"
	"log"

	"github.com/GonzaloGollo/GoPi/internal/domain"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}
	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepo(db DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("Repository create")
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("Repository get all")
	return r.db.Users, nil
}

// (r *repo) antes del nombre de la fx Define a qué tipo pertenece una función y permite acceder a los datos y comportamientos de la instancia específica.
