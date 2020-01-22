#!/usr/bin/env python3

from typing import Dict
import random 

morseCode: Dict[str, str] = {
  "A": ". _",
  "B": "_ . . .",
  "C": "_ . _ .",
  "D": "_ . .",
  "E": ".",
  "F": ". . _ .",
  "G": "_ _ .",
  "H": ". . . .",
  "I": ". .",
  "J": ". _ _ _",
  "K": "_ . _",
  "L": ". _ . .",
  "M": "_ _",
  "N": "_ .",
  "O": "_ _ _",
  "P": ". _ _ .",
  "Q": "_ _ . _",
  "R": ". _ .",
  "S": ". . .",
  "T": "_",
  "U": ". . _",
  "V": ". . . _",
  "W": ". _ _",
  "X": "_ . . _",
  "Y": "_ . _ _",
  "Z": "_ _ . .",

  "1": ". _ _ _ _",
  "2": ". . _ _ _",
  "3": ". . . _ _",
  "4": ". . . . _",
  "5": ". . . . .",
  "6": "_ . . . .",
  "7": "_ _ . . .",
  "8": "_ _ _ . .",
  "9": "_ _ _ _ .",
  "0": "_ _ _ _ _",
} # morseCode

def quiz1():
  """Randomly asks the morse code of a alphabet or digit"""
  letterDigits = list(morseCode.keys())
  for i in range(10):
    letter = random.choice(letterDigits)
    answerFormA = morseCode[letter]
    answerFormB = "".join(morseCode[letter].split())

    ans = input(f"What letter is: {answerFormA}? ")

    if ans.upper() == letter:
      print("Good!")
    else:
      print(f"Keep practicing: it was {letter}")

def main():
  print("Morse Code Test!\n")
  quiz1()


if __name__ == "__main__":
  main()
