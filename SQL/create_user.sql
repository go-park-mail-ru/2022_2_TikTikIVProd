CREATE USER ws WITH PASSWORD 'postgres_ws';

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO ws;
