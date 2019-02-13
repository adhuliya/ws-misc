#!/usr/bin/env python3
# Instanceof check vs local variable value check
class A:
  def __init__(self, x):
    self.x = x

class C(A):
  def __init__(self):
    super().__init__(10)

class D(A):
  def __init__(self):
    super().__init__(20)

class B(C):
  def __init__(self):
    super().__init__();

import time
start = end = None

b = B()

start = time.time()
for i in range(1000000):
  if isinstance(b, C):  # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

start = time.time()
for i in range(1000000):
  if b.x == 10:         # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

