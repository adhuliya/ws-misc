#!/usr/bin/env python3

nums = []
started = False
thesum = 0
themax = themin = None
thecount = 0

def getInt(msg):
    num = None
    while True:
        try:
            num = input(msg)
            if not num: 
                num = None
                break
            num = int(num)
            break
        except ValueError as ve:
            print("Enter correct integer")
        except EOFError as eof:
            num = None
            break
            print("Stream Ended")

    return num


while True:
    num = getInt("Enter Number : ")
    if num is None:
        break

    nums.append(num)
    thecount += 1
    if not started:
        themax = num
        themin = num
        thesum += num
        started = True
    else:
        if themax < num: themax = num
        if themin > num: themin = num
        thesum += num

print()
print("Numbers : ", nums)
print("Sum     : ", thesum)
print("Max     : ", themax)
print("Min     : ", themin)
if thecount:
    print("Avg     : ", thesum/thecount)


