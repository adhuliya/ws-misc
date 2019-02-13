#!/usr/bin/env python3
# comparing strings vs numbers

import time
start = end = None

x = "hellhellhellooohellohellhellhellooohello"
y = "hellhellhellooohellomellmellmellooomello"

start = time.time()
for i in range(1000000):
  if x == y:         # METHOD1
    pass
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

x = 1000000
y = 2000000
start = time.time()
for i in range(1000000):
  if x == y:         # vs METHOD2
    pass
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

