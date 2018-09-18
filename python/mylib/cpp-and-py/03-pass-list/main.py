#!/usr/bin/env python3

import asdf

if __name__ == "__main__":
    lst = list(map(float, range(10)))
    print("FROM PYTHON:", lst)
    print("Calling C function:")
    print(asdf.asdf(lst))

