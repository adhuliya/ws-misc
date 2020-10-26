-- A setup to create the default user and database.

-- role, group (group of roles), user (belongs to groups)


-- Create the database
CREATE DATABASE iol;

--------------------------------------------------------------------------------
-- START Create basic roles that will own the raw rights
--------------------------------------------------------------------------------
-- create the admin role with almost full priviledges on the pg instance
CREATE ROLE admin_role;
ALTER ROLE admin_role WITH CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'p@ssword';


CREATE ROLE ro_role;
--------------------------------------------------------------------------------
-- END   Create basic roles that will own the raw rights
--------------------------------------------------------------------------------


--------------------------------------------------------------------------------
-- START Create groups that will belong to basic roles
--------------------------------------------------------------------------------
CREATE ROLE admin_group;
GRANT admin_role TO admin_group;

CREATE ROLE ro_group;
GRANT ro_role TO ro_group;
CREATE ROLE iol_ro;
--------------------------------------------------------------------------------
-- END   Create groups that will belong to basic roles
--------------------------------------------------------------------------------


--------------------------------------------------------------------------------
-- START Create users with login previledges
--------------------------------------------------------------------------------
CREATE ROLE admin_user;
ALTER ROLE admin_user WITH LOGIN;
GRANT admin_group TO admin_user

CREATE ROLE ro_user;
ALTER ROLE ro_user WITH LOGIN;
GRANT ro_group TO ro_user;
--------------------------------------------------------------------------------
-- END   Create users with login previledges
--------------------------------------------------------------------------------


