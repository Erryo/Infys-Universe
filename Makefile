export PORT=:8000
export POSTGRES_PASSWORD=marius
export DBNAME=IU_pg_db
build: 
	@/home/infy/go/bin/templ generate
	@go build -o ./bin/app .
run:
	PORT=$(PORT) DBNAME=$(DBNAME) POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) ./bin/app
test:
	DBNAME="IU_TEST" go test -v ./... -count=1
