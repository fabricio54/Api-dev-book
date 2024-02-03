package users

import "net/http"

// cadastrar usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criar usuário"))
}

// buscar usuário
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuário"))
}

// buscar todos os usuários
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("buscar usuários"))
}

// deletando usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando usuarios"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deletando usuários"))
}

