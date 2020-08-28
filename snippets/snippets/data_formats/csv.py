#!/usr/bin/env python3

"""
How to read and write csv files using python3.
It depends on the util.py module.
"""

import csv
from typing import Iterable, Optional as Opt


from snippets import util

TMP_FILE_NAME = "tmp.csv"

def write(csvFile, content: Opt[Iterable[Iterable]]) -> None:
  """Writes content to any object that supports file API."""
  assert util.isIterable(content), f"{type(content)}"
  writer = csv.writer(csvFile, delimiter=',', quotechar='"',
                      quoting=csv.QUOTE_MINIMAL)
  if util.isIterable(content):
    for row in content:
      assert util.isIterable(row), f"{type(row)}"
      writer.writerow(row)


def _createFile(fileName: str, content: Opt[Iterable[Iterable]]) -> None:
  with open(fileName, 'w', newline='') as csvFile:
    write(csvFile, content)


def _appendFile(fileName: str, content: Opt[Iterable[Iterable]]) -> None:
  with open(fileName, 'a', newline='') as csvFile:
    write(csvFile, content)


def _modifyAge(fileName: str) -> None:
  """Modifies age of each entry.
  To modify a csv file, read each row one by one, modify it,
  write it to a new file. Then replace the
  old file with the new one.
  """
  tmpFileName = f"tmp.csv.{fileName}"
  with open(fileName, 'r', newline='') as csvFile:
    reader = csv.reader(csvFile, delimiter=',', quotechar='"')
    with open(tmpFileName, 'w', newline='') as tmpCsvFile:
      writer = csv.writer(tmpCsvFile, delimiter=',', quotechar='"',
                          quoting=csv.QUOTE_MINIMAL)
      for row in reader:
        row[1] = int(row[1]) + 1  # add 1 to age
        writer.writerow(row)

  util.removeFileIfExists(fileName)
  util.rename(tmpFileName, fileName)


def _genDummyContent():
  return [
    ["Anshuman", 35, "Gorakhpur"],
    ["Janu", 35, "Gorakhpur"],
    ["Prashant", 35, "Muradabad"],
  ]

def main():
  print(f"Removing {TMP_FILE_NAME}")
  util.removeFileIfExists(TMP_FILE_NAME)

  print(f"Creating tmp file: {TMP_FILE_NAME}")
  _createFile(TMP_FILE_NAME, _genDummyContent())

  print(f"\nAppend to tmp file: {TMP_FILE_NAME}")
  _appendFile(TMP_FILE_NAME, [["Priyanka", 30, "Varanasi"]])

  print(f"\nModify content of tmp file: {TMP_FILE_NAME}")
  _modifyAge(TMP_FILE_NAME)

  print(f"\nContent of tmp file (modified): {TMP_FILE_NAME}")
  print(util.readFromFile(TMP_FILE_NAME))

if __name__ == "__main__":
  main()

