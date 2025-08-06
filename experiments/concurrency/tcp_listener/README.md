# Instructions
1. **Run the TCP Listener**: Start the TCP listener server by executing the `go run main.go` command in the `tcp_listener` directory.
2. **Connect to the Server**: Use a TCP client to connect to the server. You can use tools like `telnet`, `netcat`.

```bash
# in case you don't have telnet installed:
# sudo apt install inetutils-telnet
$ telnet localhost 8080

> Trying 127.0.0.1...
  Connected to localhost.
  Escape character is '^]'.
  Hi there, 172.19.0.1:54312
  Connection closed by foreign host.
```

3. If you want to stop the telnet connection, you can use `Ctrl + ;:` and then write `close` to close the connection. Notice that `;:` is a key in your keyboard, Bryan Bell.

   