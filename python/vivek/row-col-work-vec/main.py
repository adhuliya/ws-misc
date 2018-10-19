#!/usr/bin/env python3

COL_FILE = "column_vector.txt"
ROW_FILE = "row_vector.txt"
WINDOW = 5  # must be odd

from collections import deque


def genColDictSeq():
  col_dict = {}  # should be a set #FIXME
  col_word_seq = []  # used for printing result
  with open(COL_FILE) as cf:
    for line in cf:
      word = line.strip()
      if word:
        col_word_seq.append(word)
        col_dict[word] = None  # dummy place-holder
  return col_dict, col_word_seq

def genRowDictSeq():
  row_dict = {}
  row_word_seq = []  # used for printing result
  with open(ROW_FILE) as rf:
    for line in rf:
      word = line.strip()
      if word:
        row_word_seq.append(word)
        row_dict[word] = {}  # init with empty dict
  return row_dict, row_word_seq


def processWindow(word_seq, row_dict, col_dict):
  """
  word_seq is a window of words represented as a deque,
  it should always be of an odd size. (e.g. 5)

  it modifies row_dict by incrementing count of neighboring words,
  if those words are present in col_dict.
  """
  middle_index = len(word_seq) // 2
  middle_word = word_seq[middle_index]

  selected_row = None

  if middle_word in row_dict:
    selected_row = row_dict[middle_word]

    for word in word_seq:
      if word in col_dict:
        if word not in selected_row:
          selected_row[word] = 0
        selected_row[word] += 1


def processText(text, row_dict, col_dict):
  """
  text is assumed to have at least one word.
  """
  padding = ["",""]
  words = []
  words.extend(padding)
  words.extend(word.strip(".,") for word in text.split())
  words.extend(padding)

  # window of size 5
  window = deque(maxlen=5)

  # setup initial window
  window.extend(words[:4])
  del words[:4]

  for word in words:
    window.append(word)
    processWindow(window, row_dict, col_dict)

  # when here, col_dict has all the counts


def print_matrix(row_dict, row_word_seq, col_word_seq):
  print()  # newline
  i = 0

  for row_word in row_word_seq:
    if i % 30 == 0:
      print("{:<14}".format("row/col"), end="")
      for col_word in col_word_seq:
        print("{: >11}".format(col_word), end="")
      print()  # newline
    i += 1

    print("{:<14}".format(row_word), end="")
    selected_row = row_dict[row_word]
    for w in col_word_seq:
      if w in selected_row:
        count = selected_row[w]
        print("{: >11}".format(count), end="")
      else:
        print("{: >11}".format(0), end="")
    print()


def main():
  """
  assuming input is standard input.
  To provide multiple files do,
    cat file1.txt file2.txt | ./main.py
  """
  row_dict, row_word_seq = genRowDictSeq()
  col_dict, col_word_seq = genColDictSeq()

  # read line by line from standard input
  while True:
    try:
      line = input().strip()
      if line:
        processText(line, row_dict, col_dict)
    except EOFError:
      break

  # after EOFError print the matrix
  print_matrix(row_dict, row_word_seq, col_word_seq)

if __name__ == "__main__":
  main()


