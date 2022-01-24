package main

import (
	"database/sql"
	"gftp/commands"
	"gftp/commands/handlers"
	"gftp/security"
	"gftp/server"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	resolver := handlers.CommandResolver{}
	db, err := sql.Open("sqlite3", "./ftp.db")

	if err != nil {
		log.Fatal(err)
	}

	authCtx := security.NewAuthContext(db)
	userProvider := security.NewSQLiteUserProvider(db)
	authService := security.NewAuthService(authCtx, userProvider)
	server.InitAuth(userProvider, authService)
	cmdProcessor := commands.NewCommandProcessor(resolver)
	server := server.NewFTPServer(2121, "localhost", cmdProcessor)
	server.Start()
}
