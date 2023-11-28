package models

//31:00
//34:00

import (
	"github.com/aprendagolang/api-pgsql/db"
)

// retorna um vetor de todo
func GetAll() (todos []Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {

		return
	}
	//fecha a conexão
	defer conn.Close()

	//retorna todos os valores e não somente um
	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {

		return
	}

	// um for para escanear as linhas que vão na sequencia
	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {

			continue
		}

		//faz a linha das informações
		todos = append(todos, todo)

	}
	return
}
