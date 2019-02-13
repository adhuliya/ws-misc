#!/usr/bin/env python3
# function call vs expression

import time
start = end = None

class A():
  def __init__(self):
    self._x = 10
    self.y = 10

  @property
  def x(self): return self._x

a = A()

start = time.time()
for i in range(1000000):
  if a.x:             # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

start = time.time()
for i in range(1000000):
  if a.y:             # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

