#!/bin/bash

docker pull postgres
docker run --name Infys-Universe -e POSTGRES_PASSWORD=marius -p 5432:5432 -d postgres
docker exec Infys-Universe createdb -U postgres IU_pg_db
docker exec Infys-Universe createdb -U postgres IU_TEST
#docker exec -ti Infys-Universe psql -U postgres
