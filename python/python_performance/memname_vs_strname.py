#!/usr/bin/env python3
# Member name retrieval vs direct name use
class A:
  def aa1(): pass
  def aa2(): pass
  def aa3(): pass
  def aa4(): pass
  def aa5(): pass
  def aa6(): pass
  def aa7(): pass
  def aa8(): pass

import time
start = end = None

a = "aa1"
z = "aa1"

start = time.time()
for i in range(10000000):
  if a == A.aa1.__name__:  # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1:", end-start)

start = time.time()
for i in range(10000000):
  if a == z:               # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2:", end-start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

