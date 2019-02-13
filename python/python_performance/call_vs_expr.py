#!/usr/bin/env python3
# function call vs expression

import time
start = end = None

def gt(x, y): return x > y

x = 10
y = 20

start = time.time()
for i in range(1000000):
  if gt(x, y):        # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

start = time.time()
for i in range(1000000):
  if x > y:             # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

