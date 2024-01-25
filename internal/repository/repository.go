package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Angstreminus/Effective-mobile-test-task/internal/apperrors"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/entity"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) CreateUser(user *entity.User) (*entity.User, apperrors.AppError) {
	user.CreatedAt = time.Now().Format("2006-02-01 15:04:05.999")
	user.IsDeleted = false
	query := `
        INSERT INTO users(name, surname, patronymic, age, gender, nationality, is_deleted, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNINIG *;
    `
	var res entity.User
	stmt, err := ur.DB.Prepare(query)
	if err != nil {
		// db op err
		return nil, nil
	}
	if err = stmt.QueryRow(
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Age,
		user.Gender,
		user.Nationality,
		user.IsDeleted,
		user.CreatedAt).Scan(&res.ID, &res.Name, &res.Surname, &res.Patronymic, &res.Age, &res.Gender, &res.Nationality, &res.IsDeleted, res.CreatedAt, res.UpdatedAt, res.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &res, nil
}

func (ur *UserRepository) EditUser(user *entity.User) (*entity.User, apperrors.AppError) {
	user.UpdatedAt = time.Now().Format("2006-02-01 15:04:05.999")
	query := `
	UPDATE users
	SET name=$1, surname=$2, patronymic=$3, age=$4, gender=$5, nationality=$6, updated_at=$7)
	WHERE id::text=$8 RETURNINIG *;`
	var res entity.User
	stmt, err := ur.DB.Prepare(query)
	if err != nil {
		// db op err
		return nil, nil
	}
	if err = stmt.QueryRow(
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Age,
		user.Gender,
		user.Nationality,
		user.UpdatedAt,
		user.ID).Scan(&res.ID, &res.Name, &res.Surname, &res.Patronymic, &res.Age, &res.Gender, &res.Nationality, &res.IsDeleted, res.CreatedAt, res.UpdatedAt, res.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &res, nil
}

func (ur *UserRepository) DeleteUser(userID uuid.UUID) apperrors.AppError {
	deletedAt := time.Now().Format("2006-02-01 15:04:05.999")
	isDeleted := true
	query := `
	UPDATE users
	SET is_deleted=$1, deleted_at=$2)
	WHERE id::text=$3;`
	stmt, err := ur.DB.Prepare(query)
	if err != nil {
		// db op err
		return nil
	}
	row, err := stmt.Exec(isDeleted, deletedAt, userID)
	if err != nil {
		return nil
	}
	rowsAff, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAff == 0 {
		return nil
	}
	return nil
}

func (ur *UserRepository) GetAllUsers(cursor string, limit int, filters map[string]string) ([]entity.User, string, error) {
	query := `SELECT * FROM users WHERE `
	args := []interface{}{}

	for key, value := range filters {
		query += fmt.Sprintf(" %s = ?", key)
		args = append(args, value)
	}

	if cursor != "" {
		query += " AND created_at < ?"
		args = append(args, cursor)
	}

	query += " ORDER BY created_at DESC LIMIT ?"
	args = append(args, limit)

	rows, err := ur.DB.Query(query, args...)
	if err != nil {
		return nil, "", err
	}
	defer rows.Close()

	users := []entity.User{}
	var lastCreatedAt string

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Patronymic, &user.Gender, &user.Nationality, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Age, &user.IsDeleted)
		if err != nil {
			return nil, "", err
		}
		users = append(users, user)
		lastCreatedAt = user.CreatedAt
	}

	return users, lastCreatedAt, nil
}
