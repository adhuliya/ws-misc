#!/usr/bin/env python3

from cw.logger import initLogger, LogLevels
initLogger(logLevel=LogLevels.DEBUG)

import logging
log = logging.getLogger(__name__)

import unittest
import sqlite_driver as sqld
import sqlite3
from sqlite_queries import queries
from datetime import datetime as dt


class TestClass(unittest.TestCase):
    def test_create_db(self):
        db = None
        try:
            db = sqld.Sqlite3()
            db.init(dbfile="sqlite.db", queries=queries)
            db.connect()
        except sqlite3.Error as e:
            log.error(e)
            self.assertTrue(False)

        rowcount, err = db.modify('1')
        if err: log.error(err); self.assertTrue(False)

        rowcount, err = db.modify('2', ("Finish testing", "asleep",
            str(dt.now())))
        if err: log.error(err); self.assertTrue(False)

        rowcount, err = db.modify('2', ("Packup and go home", "asleep",
            str(dt.now())))
        if err: log.error(err); self.assertTrue(False)


        rows, err = db.select('3')
        log.debug("DictList Rows returned:%s", rows)
        if err: log.error(err); self.assertTrue(False)


if __name__ == "__main__":
    unittest.main()


