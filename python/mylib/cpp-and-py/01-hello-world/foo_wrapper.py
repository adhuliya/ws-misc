# A wrapper module for foo.cpp program

# Note that this an unsafe example, as the
# resource memory for the Foo is leaked,
# since it is not freed up properly.
# See: https://stackoverflow.com/questions/865115/how-do-i-correctly-clean-up-a-python-object


from ctypes import cdll
lib = cdll.LoadLibrary('./libfoo.so')

def hello():
    lib.printHW()
