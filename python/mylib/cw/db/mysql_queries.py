"""
To add new queries add them here.

This module contains the queries specifically designed for sqlite which 
may not run on other databases. This module is tied with the sqlite_driver.py
module that refers this module for each query name supplied to it.

"""

"""
This is the only entity present in the file.

Don't add any other entity.
"""
queries = {
        "00001":
        """
        CREATE TABLE IF NOT EXISTS task (
            id INT PRIMARY KEY AUTO_INCREMENT,
            task VARCHAR(100),
            status ENUM ('P','C','S'), /*P:Pending, C:Complete, S:Snoozed*/
            bornon DATE,
            wakeupat DATETIME,
            naptill DATETIME
        )

        """,

        "00002":
        """
        INSERT INTO task (task, status, bornon) 
        VALUES(%s, %s, %s)
        """,

        "00003":
        """
        SELECT * FROM task LIMIT %s, %s
        """
        }

