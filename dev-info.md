# Compiling the code:
```
# In the root dir:
GOOS=linux GOARCH=amd64 go build -o creatorcoaster ./src
```
On the server, /config and /static directories go next to the binary


# Running the code in a test env:
Templ auto-compilation: (from /src)
```
templ generate -watch -proxy="http://localhost:3000" -cmd="go run ."
```

CSS reloads:
```
wgo -file=.css templ generate --notify-proxy
```

Database:
```
sudo docker compose --env-file ./src/config/.env up
```
