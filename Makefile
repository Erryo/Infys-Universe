export PORT=:8000
export POSTGRES_PASSWORD=marius
export DBNAME=IU_pg_db
build: 
	@/home/infy/go/bin/templ generate
	@go build -o ./bin/app .
run:
	PORT=$(PORT) DBNAME=$(DBNAME) POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) ./bin/app
test:
	docker exec -it Infys-Universe psql -U postgres -d postgres -c "\c IU_TEST;" -c "DROP TABLE IF EXISTS lessons Cascade;"
	docker exec -it Infys-Universe psql -U postgres -d postgres -c "\c IU_TEST;" -c "DROP TABLE IF EXISTS users Cascade;"
	docker exec -it Infys-Universe psql -U postgres -d postgres -c "\c IU_TEST;" -c "DROP TABLE IF EXISTS userlessonlink Cascade;"
	docker exec -it Infys-Universe psql -U postgres -d postgres -c "\c IU_TEST;" -c "DROP TABLE IF EXISTS usersubjectlink Cascade;"
	DBNAME="IU_TEST" go test -v ./... -count=1
