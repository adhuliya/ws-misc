README
===========

This example demonstrates the use of Postgres with Golang. The directory `db/` contains the sql files that creates the `user` table in the database (assumed to be named as `hop`).

The directory structure for `db/` is as follows:

    db
    └── postgres
        └── sql
            ├── data
            │   └── 000-t-user.sql
            └── schema
                ├── 000-f-hilo.sql
                ├── 000-t-user.sql
                ├── 001-f-deleteuser.sql
                ├── 001-f-insertuser.sql
                ├── 001-f-selectuser.sql
                └── 001-f-updateuser.sql


The files in the `db/postgres/sql/schema` are to be run on the database `hop` in the order they are lexically sorted. The first three digits decide the order in which they should run. If the first three digits of two files is the same, they can be sourced in any order.

Once the `db/postgres/sql/schema` folder is sourced, source the files in the `db/postgres/sql/data/` directory, using the same ordering rules. This will make the database ready for testing using the program.

Now simply run the program as,

    go run main1.go    #to test stored procedures

or

    go run main2.go    #to test sql queries


Why Stored Procedure?
---------------------------
It has several advantages,

1. It separates the query from the application code.
2. Database administrator can study the type of queries being used on the database without looking at the application code, and fine tune the database accordingly.
3. Queries can be tested by the database administrator for performance and other issues, independent of the code using it.
4. The queries can be rewritten to be fine tuned without changing its behavior and without needing the application to be restarted.
5. It reduces the network traffic, as applications don't have to send lengthy queries over the network.


