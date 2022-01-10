package main

func main() {
	server := NewFTPServer(21, "localhost")
	server.Start()
}
