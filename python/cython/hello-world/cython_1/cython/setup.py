#!/usr/bin/env python3

# Use command:
#   python3 setup.py build_ext --inplace;

from setuptools import setup
from Cython.Build import cythonize

setup(
    name = "Worklist Example",
    # ext_modules = cythonize("worklist.pyx"),
    ext_modules = cythonize("worklist.pyx", annotate=True),
    zip_safe = False,
)
