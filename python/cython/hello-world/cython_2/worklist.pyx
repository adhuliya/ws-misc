# worklist module


class Item:
    def __init__(self, val = 0):
        self.val = val


    def getVal(self):
        return self.val


class Worklist:
    def __init__(self):
        self.size = 0
        self.wl   = []

    def add(self, item):
        self.size += 1
        self.wl.append(item)
  
    def pop(self):
        if self.size > 0:
            self.size -= 1
            return self.wl.pop()


def main():
    worklist = Worklist()

    for i in range(4000000):
        item = Item()
        worklist.add(item)
        worklist.pop()


