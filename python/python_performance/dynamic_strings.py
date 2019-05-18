#!/usr/bin/env python3
# dynamic strings vs io.StringIO

import time
import io
start = end = None

def gt(x, y): return x > y

x = "10"
y = "20"
s = "str";

start = time.time()
for i in range(1000000):
  with io.StringIO() as sio:
    sio.write(x)
    sio.write(y)
    s = sio.getvalue()
end = time.time()
t1 = end-start
print("Total Time1: ", end - start)

start = time.time()
for i in range(1000000):
  s = f"{x}{y}"
end = time.time()
t2 = end-start
print("Total Time2: ", end - start)

print("Percent:", ((t1-t2)/t1) * 100, "% wrt first time.")

