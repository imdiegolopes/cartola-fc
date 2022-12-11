package cartola

import (
	"context"
	"database/sql"
	"dispatcher/internal/infra/db"
	"dispatcher/internal/infra/repository"
	"dispatcher/pkg/uow"
	"fmt"
)

func main() {
	fmt.Println("Starting the Cartola FC Dispatcher Service")

	ctx := context.Background()
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer dbt.Close()

	uow, err := uow.NewUow(ctx, dbt)

	if err != nil {
		panic(err)
	}
	registerRepositories(uow)
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
