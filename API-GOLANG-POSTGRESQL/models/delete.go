//38:50
package models

import "github.com/aprendagolang/api-pgsql/db"

//função que deleta as informações através de id
func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {

		return 0, err
	}
	defer conn.Close()
	//a execução é para deletar neste caso
	res, err := conn.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {

		return 0, err
	}

	//retorna as linhas afetadas, ou seja, deletadas
	return res.RowsAffected()
}
