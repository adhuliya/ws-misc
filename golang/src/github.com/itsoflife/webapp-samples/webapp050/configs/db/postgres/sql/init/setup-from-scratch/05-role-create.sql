-- login as user postgres and create user `hop`.
-- $ cat <this-file> | sudo -i -u postgres psql -f -
-- NOTE: change the role password immediately.
-- Q. HOW TO CHANGE PASSWORD MANUALLY?
-- A. Follow the following two steps,
--   $ psql postgres://hop:anshuisneo@localhost:5432/hop
--   hop=> \password

CREATE ROLE hop;

ALTER ROLE hop WITH CREATEDB;
ALTER ROLE hop WITH CREATEROLE;
ALTER ROLE hop WITH LOGIN;
ALTER ROLE hop WITH ENCRYPTED PASSWORD 'anshuisneo';
ALTER ROLE hop VALID UNTIL 'infinity';
