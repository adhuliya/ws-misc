
import worklist as wl


def main():
    worklist = wl.Worklist()

    for i in range(4000000):
        item = wl.Item()
        worklist.add(item)
        worklist.pop()


if __name__ == "__main__":
  main()



