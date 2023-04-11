package repository

import "github.com/TotrazOrstO/Imapro/pkg/database"

type UserRepository interface {
	CreateUser(user *database.User) error
	GetUserByID(id int) (*database.User, error)
}

type PostgresUserRepository struct {
	db *database.PostgresDB
}

func NewPostgresUserRepository(db *database.PostgresDB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(user *database.User) error {
	// Реализация сохранения пользователя в базу данных PostgreSQL
}

func (r *PostgresUserRepository) GetUserByID(id int) (*database.User, error) {
	// Реализация получения пользователя из базы данных PostgreSQL по ID
}
