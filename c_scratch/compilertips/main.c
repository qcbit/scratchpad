#include <stdio.h>

#ifndef TEST
#define TEST 128
#endif

int main() {
    printf("The TEST macro value is %d", TEST);
    printf("Now recompile with gcc -D'TEST=101' main.c -o main");

    return 0;
}

