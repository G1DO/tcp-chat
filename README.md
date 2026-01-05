# TCP Chat Server

A simple multi-client chat server built in Go to learn networking and concurrency.

## What it does

- Accepts multiple client connections
- Broadcasts messages from one client to all others
- Handles client disconnections gracefully

## How to run

### Start the server

```bash
go run main.go
```

### Connect clients

Open multiple terminals and run:

```bash
nc localhost 9999
```

Type a message in one terminal - it appears in all others!

## Concepts learned

| Concept | What it does |
|---------|--------------|
| `net.Listen` | Opens a port to accept connections |
| `listener.Accept` | Waits for and accepts incoming connections |
| `net.Conn` | Represents an active connection (read/write) |
| `goroutines` | Handles multiple clients concurrently |
| `sync.Mutex` | Protects shared data from race conditions |
| `defer` | Ensures cleanup happens when function exits |

## Code structure

```
main()
  └── Listen on port 9999
  └── Loop: Accept connections
        └── go handleConnection(conn)  ← runs concurrently

handleConnection(conn)
  └── addClient(conn)      ← register
  └── Loop: Read messages
        └── broadcast()    ← send to others
  └── removeClient(conn)   ← cleanup (deferred)
```

