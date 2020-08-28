#!/usr/bin/env bash

# NOTE: must be invoked as `./clean.sh` (i.e. from this directory)

# also see `.gitignore`
rm -Rf tmp.*;

cd data_formats;
../clean.sh;
cd -; # return back to current dir
