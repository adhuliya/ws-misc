#!/usr/bin/env python3

import mysql_driver as md
import unittest
import mysql.connector
import logging as log
from datetime import datetime as dt

import mysql_queries as sqr



class TestClass(unittest.TestCase):
    def test_create_db(self):
        db = None
        try:
            db = md.QueryMap(querymap=sqr.queries, passwd='anshuisneo')
        except mysql.connector.Error as err:
            log.error(e)
            self.assertTrue(False)


        rowcount, err = db.modify('00001')
        if err: log.error(err); self.assertTrue(False)

        rowcount, err = db.modify('00002', ("Finish testing", "P",
            str(dt.now())))
        if err: log.error(err); self.assertTrue(False)

        rowcount, err = db.modify('00002', ("Packup and go home", "C",
            str(dt.now())))
        if err: log.error(err); self.assertTrue(False)



        rows, err = db.select('00003', offset=0, row_count=10)
        if err: log.error(err); self.assertTrue(False)

        log.debug(rows)

        db.close()




if __name__ == "__main__":
    log.basicConfig(filename="log.txt", level=log.DEBUG,
            format=("[%(asctime)s:%(levelname)8s:%(filename)s:%(lineno)3s"
            " - %(funcName)20s() ]\n %(message)s\n"))

    unittest.main()
