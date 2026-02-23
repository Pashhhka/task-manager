package repository

import (
	"database/sql"

	"github.com/Pashhhka/task-manager/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status, user_id)
			  VALUES ($1, $2, $3, $4) RETURNING id`
	return r.DB.QueryRow(query,
		task.Title,
		task.Description,
		task.Status,
		task.UserID,
	).Scan(&task.ID)
}

func (r *TaskRepository) GetByUser(userID int) ([]models.Task, error) {
	rows, err := r.DB.Query(
		`SELECT id, title, description, status, user_id 
		 FROM tasks WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.UserID,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	query := `UPDATE tasks 
			  SET title=$1, description=$2, status=$3
			  WHERE id=$4 AND user_id=$5`
	_, err := r.DB.Exec(query,
		task.Title,
		task.Description,
		task.Status,
		task.ID,
		task.UserID,
	)
	return err
}

func (r *TaskRepository) Delete(id, userID int) error {
	_, err := r.DB.Exec(
		`DELETE FROM tasks WHERE id=$1 AND user_id=$2`,
		id,
		userID,
	)
	return err
}
