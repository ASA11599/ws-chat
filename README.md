# ws-chat

A chat server for WebSocket clients.

- Clients can list active rooms.
- Clients can connect to a room.
- Clients can send messages to the room they are connected to.
- Messages sent to a room will be broadcasted to all connected clients in that room.

## Status

[![Go](https://github.com/ASA11599/ws-chat/actions/workflows/go.yml/badge.svg)](https://github.com/ASA11599/ws-chat/actions/workflows/go.yml)

## Usage

- Run the server using the `Dockerfile`
- List active rooms: `GET /rooms` (response: `[ { "name": "<room_name>", "size": <room_size> } ]`)
- Connect and send/receive messages to/from a room:

```javascript

const ws = new WebSocket("ws://<host>:<port>/<room_name>/ws");

ws.addEventListener("message", (event) => {
    // Handle the WebSocket MessageEvent
    console.log(event);
});

// Send a message to the server
ws.send("It's lit");

// Close the connection
ws.close();

```
