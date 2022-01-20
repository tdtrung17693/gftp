package main

func main() {
	server := NewFTPServer(2121, "localhost")
	server.Start()
}
