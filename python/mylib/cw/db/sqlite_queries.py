"""
To add new queries add them here.

This module contains the queries specifically designed 
for sqlite which may not run on other databases.
This file is a sample file passes to sqlite_driver.py module,
which reads this file for each query id supplied to it.

"""

"""
This is the only entity present in the file.

Don't add any other entity.
"""
_queries = {
        "1":
        """
        CREATE TABLE IF NOT EXISTS task (
            id INT PRIMARY KEY,
            task TEXT,
            status TEXT,
            bornon TEXT,
            wakeupat TEXT,
            naptill TEXT
        )

        """,

        "2":
        """
        INSERT INTO task (task, status, bornon)
        VALUES(?, ?, ?)
        """,

        "3":
        """
        SELECT * FROM task
        """
        }

def removewhitespace(query):
    return " ".join(query.split())

def compressqueries():
    queries = dict()
    for queryid, query in _queries.items():
        queries[queryid] = removewhitespace(query)

    return queries


queries = compressqueries()


