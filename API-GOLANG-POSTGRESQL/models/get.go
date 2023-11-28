package models

import (
	"github.com/aprendagolang/api-pgsql/db"
)

// retorna um todo ou um error
func Get(id int64) (todo Todo, err error) {

	conn, err := db.OpenConnection()
	if err != nil {

		return
	}

	// fecha a conexão
	defer conn.Close()

	// retorna a linha selecionada ao id respectivo
	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	// ponteiro na ordem dos atributos criados
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	return
}
