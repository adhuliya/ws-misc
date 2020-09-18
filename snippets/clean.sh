#!/usr/bin/env bash

# NOTE: must be invoked as `./clean.sh` (i.e. from this directory)

# also see `.gitignore`
rm -Rf tmp.*;

cd snippets/data_formats || exit;
../clean.sh;
cd - || exit; # return back to current dir
