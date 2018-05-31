-- assumes the login role `hop` with appropriate permissions is already created.
-- $ psql postgres://hop:anshuisneo@localhost:5432/postgres -f <this-file>

CREATE DATABASE hop;

\c hop;
ALTER SCHEMA public OWNER TO hop
