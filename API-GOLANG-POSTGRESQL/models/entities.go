package models

//definir a struct para carregar os dados pra dentro
//do nosso banco, qd lê a gente pode tambem carregar
//os dados para poder enviar a resposta da nossa API

//definir o type TODO

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
