/*
 * Demonstration on how to define preprocessor 
 * and change its value during compilation
 *
 * We can change the following way:
 *      gcc -DTEST=101 main.c -o main
 */
#include <stdio.h>

#ifndef TEST
#define TEST 128
#endif

int main() {
    printf("The TEST macro value is %d\n", TEST);

    return 0;
}

