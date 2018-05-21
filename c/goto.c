
/*
 * goto jumps into the middle of the loop.
 * Note: this jump is not allowed in GoLang.
 * GoLang only allows jump to a parent block.
 * */
#include <stdio.h>
int main(int argc, char **argv) {
  int i = 5;
  goto THIS;

  for (i=0; i<10;i++) {
THIS:
    printf("Hello %d.\n", i);
  }

  return 0;
}
