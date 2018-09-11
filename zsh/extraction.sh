#!/usr/bin/zsh

# test script
fred='/bin/path/fred.txt.tgz';
echo ${fred:e};  # extension : tgz
echo ${fred:t};  # tail : 'fred.txt.tgz'
echo ${fred:r};  # root : '/bin/path/fred.txt'
echo ${fred:h};  # head : '/bin/path'
echo ${fred:h:h};  # head, head : '/bin'
echo ${fred:t:r};  # tail, root : 'fred.txt'
