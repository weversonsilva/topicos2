package models

import "github.com/aprendagolang/api-pgsql/db"

//essa função vai ser publica, pq vamos chamar essa função la de dentro
//dos handlers depois
//24:00
func Insert(todo Todo) (id int64, err error) {
	//abrir uma conexao com o banco
	conn, err := db.OpenConnection()
	//validar se tem erro
	if err != nil {
		return
	}
	//vai ser executado apos a função Insert for chamada e encerra o processo
	defer conn.Close()

	//criar nosso stateman, ou seja nosso SQL q vai ser executado
	//variavel sql
	//vai criar tabela, usuarios
	//Retonar um ID, pq quero fazer um insert e me retone o
	//ID depois da inserção
	sql := `INSERT INTO todos(title, description, done) VALUES ($1, $2, $3) RETURNING id`

	//mandar ele executar, vai fazer a transação com o banco de dados
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
	//28:00
}
