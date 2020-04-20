from setuptools import setup
from Cython.Build import cythonize

# How to build?
#     python3 setup.py build_ext --inplace

setup(
    ext_modules = cythonize("worklist.pyx", annotate=True)
)
