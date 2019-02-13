#!/usr/bin/env python3
# Speed test: local variable access vs member access.

class A:
  def __init__(self):
    self.c = C()

class C:
  def __init__(self):
    self.x = 10

import time
start = end = None

a = A()

start = time.time()
for i in range(1000000):
  if a.c.x == 10:         # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

x = a.c.x
start = time.time()
for i in range(1000000):
  if x == 10:           # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

