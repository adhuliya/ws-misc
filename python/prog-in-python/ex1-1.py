#!/usr/bin/env python3

import sys

zero    = ["******",
           "*    *",
           "*    *",
           "*    *",
           "*    *",
           "*    *",
           "******",]

one     = ["     *",
           "     *",
           "     *",
           "     *",
           "     *",
           "     *",
           "     *",]

two     = ["******",
           "     *",
           "     *",
           "******",
           "*     ",
           "*     ",
           "******",]

three   = ["******",
           "     *",
           "     *",
           "******",
           "     *",
           "     *",
           "******",]

four    = ["*    *",
           "*    *",
           "*    *",
           "******",
           "     *",
           "     *",
           "     *",]

five    = ["******",
           "*     ",
           "*     ",
           "******",
           "     *",
           "     *",
           "******",]

six     = ["******",
           "*     ",
           "*     ",
           "******",
           "*    *",
           "*    *",
           "******",]

seven   = ["******",
           "     *",
           "     *",
           "     *",
           "     *",
           "     *",
           "     *",]

eight   = ["******",
           "*    *",
           "*    *",
           "******",
           "*    *",
           "*    *",
           "******",]

nine    = ["******",
           "*    *",
           "*    *",
           "******",
           "     *",
           "     *",
           "******",]

digits = [zero, one, two, three, four, five, six ,seven, eight, nine]

if __name__ == "__main__":
    try:
        arg = sys.argv[1]
        num = int(arg)
        for i in range(7):
            line = ""
            for digit in arg:
                d = int(digit)
                line += digits[d][i].replace('*',digit)
                line += " "

            print(line)

    except IndexError as ie:
        print("Must give one integer as argument.")
    except ValueError as ve:
        print("Argument must be an integer.")





