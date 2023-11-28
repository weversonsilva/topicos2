//34:45
package models

import "github.com/aprendagolang/api-pgsql/db"

// recebe dois parametros (id e todo)
func Update(id int64, todo Todo) (int64, error) {

	//Abrir a conexão com o BD
	conn, err := db.OpenConnection()
	if err != nil {

		return 0, err
	}

	//Fecha a conexão
	defer conn.Close()

	//executa o update nos atributos
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, id, todo.Title, todo.Description, todo.Done)

	if err != nil {

		return 0, err
	}

	// retorna o numero de linhas que foram afetadas
	return res.RowsAffected()
} //38:20
