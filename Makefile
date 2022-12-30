.PHONY: start_bd, init_db, create_tables, drop_tables

start:
	docker-compose up -d

build-server:
	docker-compose build server 

create_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f SQL/create.sql

generate_data:
	cd gen && python3 gen.py

generate_msg_31:
	cd gen && python3 gen_msg.py 31

fill_tables: create_tables
	psql postgresql://postgres:postgres@localhost:13080/postgres -f gen/load_data.sql

fill_msg:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f gen/load_msg.sql

drop_tables:
	psql postgresql://postgres:postgres@localhost:13080/postgres -f SQL/drop_all.sql

fill_attachments:
	./attachments/download.sh

drop_attachments:
	./attachments/delete.sh

fill_files:
	./files/download.sh

drop_files:
	./files/delete.sh


