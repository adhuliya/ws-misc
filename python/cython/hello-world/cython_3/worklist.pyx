# worklist module
from libcpp.vector cimport vector


cdef class Item:
    cdef int val;

    def __init__(self, val = 0):
        self.val = val


    cdef int getVal(self):
        return self.val


cdef class Worklist:
    cdef int size
    cdef vector[Item] wl
    wl.reserve(10)

    def __init__(self):
        self.size = 0
        #self.wl   = []

    cdef void add(self, item):
        self.size += 1
        self.wl.push_back(item)
  
    cdef Item pop(self):
        if self.size > 0:
            self.size -= 1
            return self.wl.pop_back()


def main():
    worklist = Worklist()

    for i in range(4000000):
        item = Item()
        worklist.add(item)
        worklist.pop()


