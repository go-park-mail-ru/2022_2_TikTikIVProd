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

generate_data:
	cd gen && python3 gen.py

generate_msg_31:
	cd gen && python3 gen_msg.py 31

fill_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f gen/load_data.sql

fill_msg:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f gen/load_msg.sql

drop_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f SQL/drop_all.sql

fill_images:
	./images/download.sh

drop_images:
	./images/delete.sh


