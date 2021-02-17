package controllers

import "net/http"

// CreateUser method to create user and returns an User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando um usuário"))
}

// ListUsers List all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando todos os usuários"))
}

// GetUser Find one user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// UpdateUser update one user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeleteUser delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
