# The mysql driver module
import logging as log
import mysql.connector

# logger = log.getLogger(__name__)

class MySql():
    def __init__(self, config=None, passwd=None):
        self.config = {
          'user': 'root',
          # 'password': 'don't write here because this dict is logged',
          'host': '127.0.0.1',
          'port': 3306,
          'database': 'nailit',
          #'raise_on_warnings': True,
          # 'use_pure': False, # use the c implementation not python
        }
        self.conn = self.cur = None

    def init(self, config, passwd):
        if config:
            self.config = config

        try:
            self.conn = mysql.connector.connect(password=passwd, **self.config)
            self.cur = self.conn.cursor() 
            log.info("Database Conection Success. Config = %s", self.config)
        except mysql.connector.Error as err:
            log.error("Database Connection Failed. Config = %s, Error = %s", self.config, err)
            return "mysql.connector.Error"

        return None

    """
    Executes query and return if any error. If there is no error,
    the result can be extracted from self.cur object.
    This method is used internally for this class object.
    """
    def execute(self, query, data):
        try:
            if data == None:
                self.cur.execute (query)
            else:
                self.cur.execute (query, data)
            log.info("Query Success : Query = [%s], Data = [%s]", query, data)
        except mysql.connector.Error as err:
            log.error("Query Failure : Query = [%s], Data = [%s], Config = [%s], Error = [%s]", 
                    query, data, self.config, err)
            return str(err)


    """
    select query must use the suffix to the query LIMIT
    data is supplied as a tuple.
    returns the result, error.
    """
    def select (self, query, data=None, offset=0, row_count=20): 
        """
        query = ("SELECT first_name, last_name, hire_date FROM employees "
                 "WHERE hire_date BETWEEN %s AND %s")

        hire_start = datetime.date(1999, 1, 1)
        hire_end = datetime.date(1999, 12, 31)

        cursor.execute(query, (hire_start, hire_end))

        for (first_name, last_name, hire_date) in cursor:
          print("{}, {} was hired on {:%d %b %Y}".format(
            last_name, first_name, hire_date))
        """

        err = None
        if data:
            data += (offset, row_count)
        else:
            data = (offset, row_count)

        err = self.execute(query, data)

        if err:
            return None, err
        else:
            try:
                return self.cur.fetchall(), None
            except mysql.connector.Error as err:
                log.error("Query Fetchall Failure : Query = [%s], Data = [%s], Config = [%s], Error = [%s]", 
                        query, data, self.config, err)
                return None, str(err)

    """
    returns the rowcount, error
    """
    def modify (self, query, data=None):
        """
        add_salary = ("INSERT INTO salaries "
                      "(emp_no, salary, from_date, to_date) "
                      "VALUES (%(emp_no)s, %(salary)s, %(from_date)s, %(to_date)s)")


        # Insert salary information
        data_salary = {
          'emp_no': emp_no,
          'salary': 50000,
          'from_date': tomorrow,
          'to_date': date(9999, 1, 1),
        }
        
        cursor.execute(add_salary, data_salary)
        """
        err = self.execute(query, data)

        if err:
            return None, err
        else:
            try:
                self.conn.commit()
                return self.cur.rowcount, None
            except mysql.connector.Error as err:
                log.error("Query Commit Failure : Query = [%s], Data = [%s], Config = [%s], Error = [%s]", 
                        query, data, self.config, err)
                return None, str(err)



    """ Returns err string if any or else None"""
    def close (self):
        try:
            if self.cur: self.cur.close()
            if self.conn: self.conn.close()
            log.info("Database Connection Closed.")
            return None
        except:
            log.error("Database Connection NOT Closed. Config = %s", self.config)
            return str(e)

"""
If there is a query map, use this class for convenience.
"""
class QueryMap():
    def __init__(self, config=None, querymap=None, passwd=None): 
        if not config:
            self.config = {
              'user': 'root',
              # 'password': 'anshuisneo',
              'host': '127.0.0.1',
              'port': 3306,
              'database': 'nailit',
              # 'raise_on_warnings': True,
              # 'use_pure': False, # use the c implementation not python
            }
        else:
            self.config = config

        self.driver = Driver(self.config, passwd=passwd)
        self.querymap = querymap


    def select (self, queryid, data=None, offset=0, row_count=20):
        log.info("queryid = %s", queryid)

        query = self.querymap[queryid]

        return self.driver.select(query, data, offset, row_count)


    def modify (self, queryid, data=None):
        log.info("queryid = %s", queryid)

        query = self.querymap[queryid]

        return self.driver.modify(query, data)

    def close(self):
        return self.driver.close()


