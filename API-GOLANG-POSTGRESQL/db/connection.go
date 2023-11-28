//responsavel pela conexao com o banco de dados

package db

import (
	"database/sql"
	"fmt"

	//definir o driver do postgre //tem ser manual 20:50
	"github.com/aprendagolang/api-pgsql/configs"
	_ "github.com/lib/pq"
	//"golang.org/x/tools/internal/typesinternal"
)

// criar uma função que vai abrir uma conexao com o banco de dados
// no model tratar as conexões e executar as transaçoes em cima dela
func OpenConnection() (*sql.DB, error) {

	conf := configs.GetDB()
	//string de Conexao  = sc
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	//agora vai abrir a conexão com o Bando de Dados
	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}
