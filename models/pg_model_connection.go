package models

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {

	once_pg.Do(func() {
		urlString := "postgres://postgreshxh5:postgresxh5@postgres:5432/postgresxh5?pool_max_conns=45"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(context.Background(), config)
	})

	return p_pg
}
