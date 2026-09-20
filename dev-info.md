Compiling the code:
```
GOOS=linux GOARCH=amd64 go build -o creatorcoaster ./src
```
On the server, /config and /static directories go next to the binary

Running the code in a test env:
```
cd src
templ generate -watch -proxy="http://localhost:3000" -cmd="go run ./src"
# In another window:
wgo -file=.css templ generate --notify-proxy
```
