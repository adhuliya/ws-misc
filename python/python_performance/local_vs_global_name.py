#!/usr/bin/env python3
# accessing local name vs global name

import time
start = end = None

limit = 10000000

LS = True

def accessGlobal():
  for i in range(limit):
    if LS:
      pass

def accessLocal():
  LLS = LS
  for i in range(limit):
    if LLS:
      pass

start = time.time()
accessGlobal()        # METHOD1
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

start = time.time()
accessLocal()         # METHOD2
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

