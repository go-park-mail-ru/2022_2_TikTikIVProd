.PHONY: start_bd, init_db, create_tables, drop_tables

init_db:
	docker run --name ws_pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -p 13080:5432 -d postgres 

start_db:
	docker start ws_pg

stop_db: 
	docker stop ws_pg || true

rm_db: 
	docker stop ws_pg || true && docker rm ws_pg || true

create_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f SQL/create.sql

drop_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f SQL/drop_all.sql


