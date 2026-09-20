# Development

## Prerequisites:
- Go 1.26+
- [templ](https://templ.guide/quick-start/installation)
- [wgo](https://github.com/bokwoon95/wgo) (optional)

## Running locally:
1) Run the shorthand command from `./src`. This will handle automatic recompilation & reloads when you save a Go file
```bash
templ generate -watch -proxy="http://localhost:3000" -cmd="go run ."
```

2) In another terminal window, run:
```bash
wgo -file=.css templ generate --notify-proxy
```

3) This project uses MariaDB, but you can also use MySQL with no issues. The provided `docker-compose.yml` is the configuration I use for the database.
The database schema is tightly coupled to our discord bot. I recommend you to remake that part to fit your needs.
```bash
sudo docker compose --env-file ./src/config/.env up
```
Running SQL statements directly:
```
sudo docker compose --env-file ./src/config/.env exec mariadb mariadb -u root -p
```

4) `.env`: The `.env` file is in `./src/config/.env`. It looks like this:
```
INTERNAL_PORT=:3000 # The port used for this web-server

DB_USER="" # Your database user
DB_ENTRY_PORT="" # The port the database is hosted on
DB_NAME="" # Which database to use on that DB server
DB_PASSWORD="" # The password to that database

DB_DOCKER_PORT=WHATEVER:3306 # The valid docker-compose syntax for ports. The part before : is the port used to access it, and the part after is the internal port the database uses.
```

## Compiling the code:
```
# In the root dir:
GOOS=linux GOARCH=amd64 go build -o creatorcoaster ./src
```
On the server, `/config` and `/static` directories go next to the binary
