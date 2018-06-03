#!/usr/bin/env python3

import json
import getpass
import sys
import textwrap
import subprocess as sub
import os
import glob
import io


cmd1 = "sudo -i -u postgres psql -f {filename}"
cmd2 = "psql postgres://{d[username]}:{d[password]}@{d[host]}:{d[port]}/postgres -f {filename}"

options = {
        "reset" : "Purge db, setup the schema again, initialize with sample data.",
        "setup" : "Setup the db schema, and initialize with sample data. Assumes appropriate db roles already created.",
        "roles"  : "Setup roles to access database."
        }

def printHelp():
    print(textwrap.dedent("""
    Usage: ./setup.py <reset|setup|roles>

    This program manages the automatic execution of postgres
    scripts, that manage all database related operations.

    For first time setup, run the following commands,
    ./setup.py roles
    ./setup.py setup

    Currently it supports the following functions,"""))

    for key in options:
        print("{:<10} : {}".format(key, options[key]))

    exit(0)
        
def sortedSqlFileList(directory):
    """
    Assumes directory string end with forward slash.
    """
    files = glob.glob(directory + "*.sql")
    files.sort()
    return files

def process(conf, dirpath):
    """
    Sources the *.sql files in the given dir.
    """
    error = False
    files = sortedSqlFileList(dirpath)
    print("hello", files)
    for fname in files:
        cp = sub.run(cmd2.format(d=conf, filename=fname), shell=True)

        if cp.returncode == 0:
            print("Sourced:", fname)
        else:
            error = True
            print("Error  :", fname, cp.returncode, file=sys.stderr)

    if error:
        print("FINISHED WITH ERRORS.", file=sys.stderr)
        return False
    else:
        return True

def setupRoles(conf, dirpath):
    """
    Creates roles to access database.
    """
    files = sortedSqlFileList(dirpath)
    alltext = io.StringIO()
    for fname in files:
        print("Sourcing:", fname)
        with open(fname) as f:
            alltext.write(f.read())

    tmpFile = sub.getoutput("mktemp")
    print("TempFile:", tmpFile)
    cp = sub.run("chmod +r {}".format(tmpFile), shell=True)
    if cp.returncode != 0:
        print("Error changing permission of tempfile {}".format(tmpFile))
        return False # no use continuing

    with open(tmpFile, "w") as f:
        f.write(alltext.getvalue())

    print("Enter UNIX password if prompted.")
    error = False
    cp = sub.run(cmd1.format(filename=tmpFile), shell=True)

    if cp.returncode == 0:
        print("Sourced:", tmpFile)
    else:
        error = True
        print("ERROR:", fname, cp.returncode, file=sys.stderr)

    sub.run("rm {}".format(tmpFile), shell=True)

    if error:
        print("FINISHED WITH ERRORS.", file=sys.stderr)
        return False
    else:
        return True

def readConfigs(filename):
    with open(filename) as f:
        conf = json.loads(f.read())
        return conf

if __name__ == "__main__":
    if len(sys.argv) > 1:
        cmd = sys.argv[1].strip().lower()
    else:
        printHelp()

    if cmd in options:
        conf = readConfigs("config.json")

    good = True
    if cmd == "reset":
        conf["password"] = getpass.getpass("postgres password: ")
        good = good and process(conf, "./sql/init/reset/")
        good = good and process(conf, "./sql/schema/")
        good = good and process(conf, "./sql/data/")
    elif cmd == "setup":
        conf["password"] = getpass.getpass("postgres password: ")
        good = good and process(conf, "./sql/init/setup/")
        good = good and process(conf, "./sql/schema/")
        good = good and process(conf, "./sql/data/")
    elif cmd == "roles":
        good = setupRoles(conf, "./sql/init/setup-roles/")
    else:
        printHelp()

    if good: print("SUCCESS!")


