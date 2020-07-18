-- A setup to create the default user and database.

CREATE ROLE itsoflife;
ALTER ROLE itsoflife WITH CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'anshuisneo';
SET ROLE itsoflife;
CREATE DATABASE itsoflife; -- db with same name as user's is a convenience.
