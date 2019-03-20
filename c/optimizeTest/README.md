`optimizeTest.c` has a program that should have been optimized by GCC, Clang, and Lerner's approach. See the generated assembly files (with `-O3` flag) for confirmation.

SPAN Framework can workout that the else condition of the `if` statement is not feasible.


`optimizeTest1.c` is a more simpler example where SPAN works and GCC, Clang and Lerner's approach fail to recognize the dead code.

