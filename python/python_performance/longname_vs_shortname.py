#!/usr/bin/env python3
# Long name access vs sort name access

import time
start = end = None

a_very_long_name_for_variable = 10

start = time.time()
for i in range(1000000):
  if a_very_long_name_for_variable == 10:   # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

a = 10
start = time.time()
for i in range(1000000):
  if a == 10:           # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

