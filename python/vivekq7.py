#!/usr/bin/env python3

# assuming standard input is used for the given file
# correct input format with no blank or extra lines is assumed.
# leading/trailing spaces are automatically removed

from datetime import datetime as dt
from datetime import timedelta
from collections import OrderedDict
import re

itemPattern = re.compile("\s*(?P<start>\S{5})\s+(?P<end>\S{5}).*")

def parseItem(itemstr):
    m = itemPattern.match(itemstr.strip())
    start = dt.strptime(m.group("start"), "%H:%M")
    stop  = dt.strptime(m.group("end"), "%H:%M")

    return (start, stop)


def readInput():
    schedule = OrderedDict()
    while True:
        try:
            datestr = input().strip()
            schedule[datestr] = []
            num = int(input()) # num of items
            for i in range(num):
                item = parseItem(input())
                schedule[datestr].append(item)
            # assuming that input schedule may not be sorted
            schedule[datestr].sort(key = lambda x : x[0])
        except EOFError as e:
            break

    return schedule

def findSlots(schedule):
    slots = OrderedDict()
    day_begin = dt.strptime("09:00", "%H:%M")
    day_end   = dt.strptime("18:00", "%H:%M")

    for key, val in schedule.items():
        slots[key] = None # allocate a place in the dict

        # find first max duration
        maxDuration = timedelta(0)
        maxDurStart = dt.strptime("09:00", "%H:%M")
        maxDurEnd   = dt.strptime("09:00", "%H:%M")

        begin = day_begin
        for (start, end) in val:
            # check free time duration
            if maxDuration < start - begin:
                maxDuration = start - begin
                maxDurStart = begin
                maxDurEnd   = start
            begin = end # the beginning of the next possible free time

        # check the last possible slot
        if maxDuration < day_end - begin:
            maxDuration = day_end - begin
            maxDurStart = begin
            maxDurEnd   = day_end

        slots[key] = (maxDurStart, maxDurEnd)

    return slots


def printSlots(slots):
    # using the array to maintain order
    for key, val in slots.items():
        print(key, (val[1] - val[0]).seconds//60, val[0].strftime("%H:%M"), val[1].strftime("%H:%M"))

    
def process():
    printSlots(findSlots(readInput()))


if __name__ == "__main__":
    process()


