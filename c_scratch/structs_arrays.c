#include <stdio.h>
#include <assert.h>

typedef struct {
	int a, b;
	double c, d;
	int *efg;
} demo_s;

int main(void){
	demo_s d1 = {.b=1, .c=2, .d=3, .efg=(int[]){4,5,6}};
	demo_s d2 = d1;

	printf("d1.a=%d\n", d1.a);
	printf("d2.a=%d\n", d2.a);

	printf("Setting d1.b=14\n");
	d1.b = 14;
	printf("d2.b=%d\n", d2.b);
	printf("Setting d1.c=41\n");
	d1.c = 41;
	printf("d2.c=%f\n", d2.c);
	printf("d2.efg[0]=%d\n", d2.efg[0]);
	printf("Setting d1.efg[0]=7\n");
	d1.efg[0]=7;
	printf("d2.efg[0]=%d\n", d2.efg[0]);
	printf("d2.efg's value changed because declaring demo_s d2 = d1 copied the pointer, which points to the original d1.efg still.\n");

	printf("d2.b=%d\n", d2.b);

	printf("Assert - is to state a fact confidently and completely.\n");
	printf("So when we assert(d2.b==1) then program will not exit, because d2.b = 1.\n");
	assert(d2.a==0);
	assert(d2.b==1);
	printf("But what if we assert(d2.b==0)? The program will exit, because this is not a fact, i.e. it is not true.\n.");
	//assert(d2.b==0);
	assert(d2.c==2);
	assert(d2.d==3);
	assert(d2.efg[0]==7);

	return 0;
};
