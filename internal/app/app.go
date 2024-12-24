package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Task interface defines the methods for a task.
type Task interface {
	GetID() int
	GetDescription() string
	GetStatus() string
	SetDescription(description string)
	SetStatus(status string)
	SetUpdatedAt(updatedAt time.Time)
}

// TaskImpl is a concrete implementation of the Task interface.
type TaskImpl struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (t *TaskImpl) GetID() int {
	return t.ID
}

func (t *TaskImpl) GetDescription() string {
	return t.Description
}

func (t *TaskImpl) GetStatus() string {
	return t.Status
}

func (t *TaskImpl) SetDescription(description string) {
	t.Description = description
}

func (t *TaskImpl) SetStatus(status string) {
	t.Status = status
}
func (t *TaskImpl) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

// TaskManager interface defines methods for managing tasks.
type TaskManager interface {
	AddTask(description string)
	UpdateTask(id int, newDesc string)
	DeleteTask(id int)
	ListTasks()
	ListFillterByStatus(status string)
	MarkTaskInProgress(id int)
	MarkTaskDone(id int)
	LoadTasks() error
	SaveTasks() error
}

// TaskManagerImpl is a concrete implementation of the TaskManager interface.
type TaskManagerImpl struct {
	filePath string
	tasks    []Task
	nextID   int
}

// NewTaskManager is a factory function for creating a TaskManager.
func NewTaskManager(filePath string) (*TaskManagerImpl, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("error getting absolute path: %w", err)
	}
	tm := &TaskManagerImpl{
		filePath: absPath,
		tasks:    make([]Task, 0),
		nextID:   1,
	}
	err = tm.LoadTasks()
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (tm *TaskManagerImpl) LoadTasks() error {
	data, err := os.ReadFile(tm.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("tasks.json not found, creating a new one.")
			tm.nextID = 1
			return nil // No error if file doesn't exist
		}
		return fmt.Errorf("error reading tasks file: %w", err)
	}

	var tempTasks []*TaskImpl
	err = json.Unmarshal(data, &tempTasks)
	if err != nil {
		return fmt.Errorf("error unmarshalling tasks: %w", err)
	}

	tm.tasks = make([]Task, len(tempTasks))
	for i, task := range tempTasks {
		tm.tasks[i] = task
	}

	if len(tm.tasks) > 0 {
		tm.nextID = tm.tasks[len(tm.tasks)-1].GetID() + 1
	} else {
		tm.nextID = 1
	}
	return nil
}

func (tm *TaskManagerImpl) SaveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks: %w", err) // Wrap the error for better context
	}
	return os.WriteFile(tm.filePath, data, 0644)
}

func (tm *TaskManagerImpl) AddTask(description string) {
	if description == "" {
		fmt.Println("Description cannot be empty")
		return
	}
	task := &TaskImpl{ID: tm.nextID, Description: description, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	tm.tasks = append(tm.tasks, task)
	if err := tm.SaveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err) //More informative error message
	}
	fmt.Printf("Task added successfully (ID: %d)", tm.nextID)
	tm.nextID++
}

func (tm *TaskManagerImpl) findTask(id int) (Task, error) {
	for _, task := range tm.tasks {
		if task.GetID() == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", id)
}

func (tm *TaskManagerImpl) UpdateTask(id int, newDesc string) {
	task, err := tm.findTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if newDesc != "" {
		task.SetDescription(newDesc)
		task.SetUpdatedAt(time.Now())

	}
	if err := tm.SaveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err) //More informative error message
	}
	fmt.Println("Task updated successfully")
}

func (tm *TaskManagerImpl) DeleteTask(id int) {
	index := -1
	for i, task := range tm.tasks {
		if task.GetID() == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}
	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
	if err := tm.SaveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err) //More informative error message
	}
	fmt.Println("Task deleted successfully")
}

func (tm *TaskManagerImpl) ListTasks() {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tm.tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.GetID(), task.GetDescription(), task.GetStatus())
	}
}

func (tm *TaskManagerImpl) ListFilterByStatus(status string) {
	if len(tm.tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	filteredTasks := []Task{}
	for _, task := range tm.tasks {
		if task.GetStatus() == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	if len(filteredTasks) == 0 {
		fmt.Printf("No tasks found with status '%s'.\n", status)
		return
	}
	for _, task := range filteredTasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s\n", task.GetID(), task.GetDescription(), task.GetStatus())
	}
}

func (tm *TaskManagerImpl) MarkTaskInProgress(id int) {
	task, err := tm.findTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	task.SetStatus("in-progress")
	if err := tm.SaveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err) //More informative error message
	}
	fmt.Println("Task marked as in progress successfully")
}

func (tm *TaskManagerImpl) MarkTaskDone(id int) {
	task, err := tm.findTask(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	task.SetStatus("done")
	if err := tm.SaveTasks(); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err) //More informative error message
	}
	fmt.Println("Task marked as done successfully")
}
