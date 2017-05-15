#!/usr/bin/env python3

import base64
from io import StringIO


def normalize(thestr):
    """

    Removes leadin/trailing spaces, and converts to lower case

    >>> a = normalize("  \t123AsdfG*    ")

    >>> normalize(a)
    '123asdfg*'

    using explicit blanklines :
    >>> print("\\n\\nhello")
    <BLANKLINE>
    <BLANKLINE>
    hello

    """
    return thestr.strip().lower()


def b64encode(databytes):
    """
    Bytes --to-> custom base64 String

    >>> b64encode(b'1234')
    'MTIzNA..'
    >>> b64encode(b'\\xb5\\xeb-\\x8ax>\\xfd\\xfd')
    'testing-_f0.'
    """

    # default encoding
    b64_default = base64.b64encode(databytes)

    # convert to a string
    b64_default = b64_default.decode("utf-8")

    b64_custom = StringIO()

    for ch in b64_default:
        if ch == "+":
            b64_custom.write("-")
        elif ch == "/":
            b64_custom.write("_")
        elif ch == "=":
            b64_custom.write(".")
        else:
            b64_custom.write(ch)

    val = b64_custom.getvalue()
    b64_custom.close()

    return val


def b64decode(b64_custom):
    """
    b64 string (custom or default) --to-> bytes data

    >>> b64decode("MTIzNA==")
    b'1234'
    >>> b64decode("MTIzNA..")
    b'1234'
    >>> b64decode("testing+/f0")
    b'\\xb5\\xeb-\\x8ax>\\xfd\\xfd'
    >>> b64decode("testing+_f0.")
    b'\\xb5\\xeb-\\x8ax>\\xfd\\xfd'
    """

    count = 0
    b64_default = StringIO()

    for ch in b64_custom:
        count += 1
        if ch == "-":
            b64_default.write("+")
        elif ch == "_":
            b64_default.write("/")
        elif ch == ".":
            b64_default.write("=")
        else:
            b64_default.write(ch)

    while count % 4 != 0:
        b64_default.write("=")
        count += 1

    val = base64.b64decode(b64_default.getvalue())
    b64_default.close()

    return val


if __name__ == "__main__":
    """
    Only to test this module.
    You can write doctests in a separate file too.
    Refer: http://pythontesting.net/framework/doctest/doctest-introduction/

    One can also use:
    python3 -m doctest String.py
    python3 -m doctest -v String.py #for verbosity

    """
    import doctest
    doctest.testmod()


