package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RhoanBarioni/GoLang_Tutorial/internal/jsonutil"
	_ "github.com/go-sql-driver/mysql"
)

/*

Dentro da aplicação, você ira modificar essas 3 operações:

Criação,
Leitura,
Atualização,
Remoção,

Quando um aluno for informado, você deve consultar no banco de dados para saber se ele já existe, se não, você deve cria-lo utilizando o INSERT, se não, você deve atualiza-lo utilizando o "UPDATE".

Na sequência do cenário de remoção, você deve utilizar o comando "DELETE" para apagar um aluno do banco.

E por fim, faça um SELECT trazendo todos os alunos do banco, e exibindo a lista no final.

*/

func DbNew() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/escola")

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAlunos(db *sql.DB) ([]jsonutil.Aluno, error) {
	// usar o * para modificar diretamente o valor na memória
	// usar o pacote de json só para usar a CLASSE ALUNOS
	alunos := []jsonutil.Aluno{}
	// atribuir var com a classe ALUNOS

	rows, err := db.Query("SELECT * FROM alunos")
	// usar o QUERY para fazer consultas no banco

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		aluno := jsonutil.Aluno{}
		rows.Scan(&aluno.Id, &aluno.Nome, &aluno.Media)
		alunos = append(alunos, aluno)
	}

	defer rows.Close()

	return alunos, nil
}

func GetAlunoByName(db *sql.DB, name string) (jsonutil.Aluno, error) {
	aluno := jsonutil.Aluno{}

	err := db.QueryRow("select * from alunos where nome = ?", name).Scan(
		&aluno.Id,
		&aluno.Nome,
		&aluno.Media,
	)

	if err != sql.ErrNoRows {
		return aluno, err
	}

	return aluno, nil
}

// A gente faz retornar error pq nós não iremos precisar da informação de alunos no main, só queremos saber se tudo ocorreu bem, por isso tem o err

func CreateAluno(db *sql.DB, aluno *jsonutil.Aluno) error {
	// nao preciso saber do db pq n quero q ele pingue o resultado lá fora, so quero saber se tem erro
	_, err := db.Exec("insert into alunos (nome, media) values (?, ?)", aluno.Nome, aluno.Media)
	fmt.Println("Usuario inserido com sucesso")

	if err != nil {
		return err
	}

	return nil
}

func UpdateAluno(db *sql.DB, aluno *jsonutil.Aluno) error {
	_, err := db.Exec("update alunos set nome = ?, media = ? where nome = ?", aluno.Nome, aluno.Media, aluno.Nome)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAluno(db *sql.DB, nome string) error {
	_, err := db.Exec("delete from alunos where nome = ?", nome)

	if err != nil {
		return err
	}

	return nil
}

func SalvarAluno(db *sql.DB, aluno *jsonutil.Aluno) error {
	res, err := GetAlunoByName(db, aluno.Nome)
	if err != nil {
		return err
	}

	aluno.Media = Calc(aluno.Nome, aluno.Notas)

	fmt.Println(aluno.Media)
	fmt.Println(aluno)

	if res.Nome != aluno.Nome {
		err = CreateAluno(db, aluno)
		if err != nil {
			return err
		}
	} else {
		log.Println(aluno.Media)
		err = UpdateAluno(db, aluno)
		log.Println(aluno.Media)
		if err != nil {
			return err
		}
	}

	return nil
}