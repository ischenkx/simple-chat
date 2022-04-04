The easiest way to send requests to the server is to run
the cli tool provided in `pkg/http/cli`

1. `go run pkg/http/cli/main.go <URL>`
2. Checkout the available commands below:

- `seturl` - set the server address
- `info` - current user info
- `login`
- `logout`
- `register`
- `chats` - get user chats
- `friends` - get user friends
- `incoming-fr` - get user incoming friend requests 
- `outgoing-fr` - get user outgoing friend requests 
- `send-fr` - send a friend request
- `accept-fr` - accept a friend request
- `decline-fr` - decline a friend request
- `chat` - get chat info
- `delete-chat`
- `create-chat`
- `create-chat-member` - add a member to a chat
- `delete-chat-member` - remove a member from a chat
- `get-chat-members`
- `send-message`
- `update-message`
- `delete-message`
- `messages` - get user messages from a specified chat
- `kill` - stop the process