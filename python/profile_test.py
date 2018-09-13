#!/usr/bin/env python3

def test():
    lst = list(range(100000))
    lst.sort(key=lambda x: -x)
    return lst

if __name__ == "__main__":
    lst = test()
    print(lst)
