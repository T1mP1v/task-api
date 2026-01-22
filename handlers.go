package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetTasks godoc
// @Summary Получить список задач
// @Description Получение всех задач с возможной фильтрацией
// @Tags tasks
// @Accept json
// @Produce json
// @Param status query string false "Статус задачи"
// @Success 200 {array} Task
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	status := c.Query("status")

	query := "SELECT id, title, completed FROM tasks"
	if status == "completed" {
		query += " WHERE completed = 1"
	}

	rows, err := DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.Completed)
		tasks = append(tasks, t)
	}

	c.JSON(http.StatusOK, tasks)
}

// GET /tasks/:id
func GetTask(c *gin.Context) {
	id := c.Param("id")
	row := DB.QueryRow("SELECT id, title, completed FROM tasks WHERE id = ?", id)

	var t Task
	err := row.Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	c.JSON(http.StatusOK, t)
}

// CreateTask godoc
// @Summary Создать задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "Новая задача"
// @Success 201 {object} Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var t Task
	if err := c.BindJSON(&t); err != nil || t.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	result, _ := DB.Exec(
		"INSERT INTO tasks (title, completed) VALUES (?, ?)",
		t.Title, t.Completed,
	)

	id, _ := result.LastInsertId()
	t.ID = int(id)

	c.JSON(http.StatusCreated, t)
}

// UpdateTask godoc
// @Summary Обновить задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body Task true "Обновлённая задача"
// @Success 200 {object} Task
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t Task
	c.BindJSON(&t)

	_, err := DB.Exec(
		"UPDATE tasks SET title=?, completed=? WHERE id=?",
		t.Title, t.Completed, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Задача обновлена"})
}

// DeleteTask godoc
// @Summary Удалить задачу
// @Tags tasks
// @Param id path int true "ID задачи"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
