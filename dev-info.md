templ generate -watch -proxy="http://localhost:3000" -cmd="go run ./src"

wgo -file=.css templ generate --notify-proxy