package main

import (
	"database/sql"
	"flag"
	"gftp/commands"
	"gftp/commands/handlers"
	"gftp/security"
	"gftp/server"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dataRoot = flag.String("data-root", "/media/Data1/Code/gftp-working-dir", "The root directory for storing data")
var host = flag.String("host", "0.0.0.0", "Host to which the server will bind")
var port = flag.Int("port", 21, "Port on which the server will listen")

func main() {
	flag.Parse()
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

	server := server.NewFTPServer(*port, *host, *dataRoot, cmdProcessor)

	server.Start()
}
