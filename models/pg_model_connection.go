package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {
	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg.Do(func() {
		urlString := "postgres://postgreshxh5:dfsdf4FERg45234SERFsdrf346erbeg@postgres:5432/postgresxh5?pool_max_conns=45"
		config, _ := pgxpool.ParseConfig(urlString)
		p_pg, _ = pgxpool.ConnectConfig(ctx, config)
	})

	return p_pg
}
