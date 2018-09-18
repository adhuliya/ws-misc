#!/usr/bin/env python3
print("hello")
from distutils.core import setup, Extension

module1 = Extension('asdf', sources = ['asdf_module.c'])

setup (name = 'asdf',
        version = '1.0',
        description = 'This is a demo package',
        ext_modules = [module1])
