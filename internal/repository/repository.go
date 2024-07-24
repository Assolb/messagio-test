package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"messagio/internal/config"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

type Repository struct {
	MessageRepository *MessageRepository
}

func NewDatabase(config *config.DatabaseConfig) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logrus.Println("Database is successfully connected!")

	return &Database{Db: db}, nil
}

func NewRepository(db *Database) (*Repository, error) {
	repo := &Repository{
		MessageRepository: NewMessageRepository(db),
	}

	if err := repo.MessageRepository.CreateTable(); err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return repo, nil
}
