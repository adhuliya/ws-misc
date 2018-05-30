postgres database
========================

Create database `hop`
------------------------
In order to source the sql files in the `sql/` folder, the database `hop` must first be created. Once created, the user can source the files in the `sql/schema/` folder first, followed by the `sql/data/` folder.

The sql files are deliberately ordered such that their lexical sequence doesn't cause any dependence problem.
