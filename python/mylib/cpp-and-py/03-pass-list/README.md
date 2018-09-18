Calling C++ from Python
========================

STILL UNDER PROGRESS. CODE DOESNOT WORK.

Use [this][1] and consequently [this][2] resources to communicate between python and C++ and taking care of the freeing up of any system resources hence created.

The `main.py` calls the C++ functions using the module `foo_wrapper.py`. The source files are documented enough to replicate the process for any other similar scenario.

For convenience, use `./build_so.sh foo.cpp` command to build `libfoo.so` file.


[1]: https://stackoverflow.com/questions/145270/calling-c-c-from-python "calling c function from python"
[2]: https://stackoverflow.com/questions/865115/how-do-i-correctly-clean-up-a-python-object "method to free a system resource"
