#include <iostream>
#include <cstdlib>
#include <time.h>

#define MAXLEN 100

struct stack_t {
    int sp;
    int stack[MAXLEN];
};

struct element_pair {
    int status;
    int element;
};

stack_t* newStack() {
    stack_t *stk = new stack_t{};
    stk->sp = 0;

    return stk;
}

void deleteStack(stack_t *stk) {
    delete stk;
}

int push(stack_t *stk, int elem) {
    if (stk->sp == MAXLEN) {
        return 0; // false
    }

    stk->stack[stk->sp] = elem;
    ++(stk->sp);
    return 1; // true
}

element_pair pop(stack_t *stk) {
    element_pair ep = element_pair{};

    if (stk->sp == 0) {
        ep.status = 0; // false
    } else {
        ep.element = stk->stack[stk->sp - 1];
        ep.status = 1; // true

        --(stk->sp);
    }

    return ep;
}

void printStack(stack_t *stk) {
    int i = 0;
    for (i = 0; i < stk->sp; ++i) {
        std::cout << stk->stack[i] << " ";
    }
    std::cout << "\nTotal: " << stk->sp << ".\n";
}

int main(int argc, char **argv) {
    stack_t *stk = newStack();

    srand(time(NULL));

    //int randVal = rand() % MAXLEN;
    int randVal = 300;
    for (int i=0; i < randVal; ++i) {
        push(stk, rand() % 100);
    }

    printStack(stk);
    deleteStack(stk);

    return 0;
}
