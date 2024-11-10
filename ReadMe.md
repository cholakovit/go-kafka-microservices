### How to create initial project:

===
mkdir producer & cd producer
go mod init producer

===
mkdir worker & cd producer
go mod init producer


### How to run the project:

main folder where docker-compose.yml is located
docker-compose up -d

===
cd /producer
go run main.go

===
cd /worker
go run main.go


### CompileDaemon is a tool that watches your Go files for changes and automatically rebuilds and restarts your Go application whenever you save changes.

cd /producer
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -command="./producer" -directory="./"

===
cd /worker
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -command="./worker" -directory="./"

### POST A MESSAGE TO PRODUCER
POST: http://localhost:3000/api/v1/comments
{ "text":"message 1" }
