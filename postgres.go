package main

import (
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type UserRec struct {
	User string
	DisplayName string
	ID int
}

func GetUser(ctx context.Context, conn *sql.DB) (UserRec, error) {
	const query = SELECT "User", "DisplayName" FROM users WHERE "ID" = $1
	u := UserRec(ID: id)
	err := conn.QueryRowContext(ctx, query, id).Scan(&u)
	return u, err
}

func main() {
	dbURL := "postgres://admin:root@localhost:5432/go_db"
	conn, err := sql.Open("pgx", dbURL)
	if err != nil {
        fmt.Printf("connection error: %s\n", err)
		return
    }
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := conn.PingContext(ctx); err != nil {
		return err
	}
	cancel()
}