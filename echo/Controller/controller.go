package Controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"echo/Model" // Update with your actual package name

	"github.com/labstack/echo/v4"
)

var db *sql.DB // Initialize your database connection in main.go and pass it to the controller

// GetUsers retrieves all users
func GetUsers(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching users")
	}
	defer rows.Close()

	var users []*Model.User
	for rows.Next() {
		user := new(Model.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return c.JSON(http.StatusInternalServerError, "Error scanning users")
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user Model.User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	user := new(Model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating user")
	}

	userID, _ := result.LastInsertId()
	user.ID = int(userID)
	return c.JSON(http.StatusCreated, user)
}

// UpdateUser updates an existing user
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user := new(Model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	_, err = db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating user")
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting user")
	}

	return c.JSON(http.StatusOK, "User deleted")
}
