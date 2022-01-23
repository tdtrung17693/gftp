# Overview

```
┌──────────────────┐                    
│server            │                    
└┬────────────────┬┘                    
┌▽──────────────┐┌▽────────────────────┐
│command process││data transfer process│
└───────────────┘└─────────────────────┘
```

The server consist of 2 process: command process and data transfer process (DTP). Every client 
connects to the server will have a `ConnContext` which hold necessary information about the client
such as authentication state, user's root folder, etc.. as well as message buses to communicate with
DTP and Command process

# Command process
Command process is the main server process to handle client's commands and interact with data
transfer process in order to transfer data back and forth. It will listen on the default port 
as described in the specs (21) if there isn't any port provided.

To parse and handle the commands from clients, in the current implementation, it will read the request
line by line and then pass the whole line to a command processor, which is compatible with the 
`CommandProcessor` interface. The command processor will then parse and make call to a registered command 
handler. All the command handlers must be compatible with the `CommandHandler` interface.

# Data transfer Process
Data transfer process is responsible for managing all the data transfer connection. When a data transfer
request is made, it will be forward to this process and will be handled there. DTP is implemented using
goroutine and all other components will interface with it using message channels.
