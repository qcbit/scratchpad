#include <stdio.h>
#include <stdlib.h>

int main(void){
	printf("int *i = malloc(sizeof(int)); //right\n");
	int *i = malloc(sizeof(int));
	printf("*i = 23; //right\n");
	*i = 23;
	printf("int *j = 23; //wrong\nwill give a complier warning\n");
	printf("warning: incompatible integer to pointer conversion initializing\n");
	int *j = 23;
	printf("It is wrong.");
	printf("Use the following rul.\n");
	printf("When used for a declaration, a star indicates a pointer\n");
	printf("When not used as a declaration, a star indicates the value of the pointer\n");
};
