#!/usr/bin/env python3
"""
 This program reads a dictionary file (one word per line) and
performs search on the words stored as List and Dict,
and compares the search time.

It selects some words randomly from the word list to search
the list.

The speedup observed is nearly 5000x to 9000x !!!
"""

import random as rand
import time

fileName = "american-english"

def readWords(fileName):
    lst = []
    with open(fileName) as f:
        for word in f:
            lst.append(word[:-1])

    return lst

def searchList(selectedWords, wordList):
    """
    search each word in `selectedWords`,
    in the list `wordList`.
    """
    for word in selectedWords:
        if word in wordList:
            pass

def searchDict(selectedWords, wordDict):
    """
    search each word in `selectedWords`,
    in the dict `wordDict`.
    """
    for word in selectedWords:
        if word in wordDict:
            pass

def timeIt(func, arg1, arg2):
    start = time.process_time()
    func(arg1, arg2)
    end = time.process_time()

    print(func.__name__ + "'s duration is", end - start)

    return end - start



def main():
    wordList = readWords(fileName)
    #print(wordList)

    # shuffle the list randomly in-place
    rand.shuffle(wordList)

    # select first few words
    selectedWords = wordList[:500]

    # reshuffle the list
    rand.shuffle(wordList)


    t1 = timeIt(searchList, selectedWords, wordList)
    
    t2 = timeIt(searchDict, selectedWords, dict((word,True) for word in wordList))


    print("Speedup :", t1/t2)


if __name__ == "__main__":
    main()
