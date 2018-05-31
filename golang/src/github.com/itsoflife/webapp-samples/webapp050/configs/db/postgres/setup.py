#!/usr/bin/env python3

import json
import getpass
import sys

cmd = "psql postgres://{d['username']}:{d['password']}@{d['host']}:{d['port']}/{d['dbname']} -f {filename}"

def readConfigs(filename):
    conf = json.loads(open(filename).read())
    return conf

def setup():
    pass

if __name__ == "__main__":
    if len(sys.argv) > 1:
        cmd = sys.argv[1]

    conf = readConfigs("config.json")
    conf["password"] = getpass.getpass("Password:")

    print(conf)
    setup()
