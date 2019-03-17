#include <stdio.h>
////
// A pointer to a block of memory that has already been automatically freed is
// worse than useless.
////

typedef struct powers {
	double base, square, cube;
} powers;

powers get_power(double in){
	powers out = {.base   = in,
		      .square = in * in,
		      .cube   = in * in * in};
	return out;
}

int *get_even(int count){
	int out[count];
	for (int i=0; i < count; i++)
		out[i] = 2*i;
	return out; //bad
}

int main() {
	powers threes = get_power(3);
	int *evens = get_even(3);
	printf("threes: %g\t%g\t%g\n", threes.base, threes.square, threes.cube);
	printf("This next one is bad, because the array has already been destroyed.\n");
	printf("evens: %i\t%i\t%i\n", evens[0], evens[1], evens[2]);	
	printf("Use memmove to copy an array.\n See memmove.c\n");

};
