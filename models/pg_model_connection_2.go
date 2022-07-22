package models

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgconn"
)

var PostgresCN_2 = Conectar_Pg_DB_2()

var (
	once_pg_2 sync.Once
	p_pg_2    *pgconn.PgConn
)

func Conectar_Pg_DB_2() *pgconn.PgConn {
	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	once_pg_2.Do(func() {
		urlString := "postgres://postgreshxh5:dfsdf4FERg45234SERFsdrf346erbeg@postgres:5432/postgresxh5"
		config, _ := pgconn.ParseConfig(urlString)
		p_pg_2, _ = pgconn.ConnectConfig(ctx, config)
	})

	return p_pg_2
}
