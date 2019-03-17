#include <assert.h>
#include <stdio.h>
#include <string.h> // memmove

int main(){
	printf("Demonstrating copying an array using memmove.\n");
	printf("int abc[] = {0,1,2};\n");
	int abc[] = {0,1,2};
	printf("int *copy1, copy2[3];\n");
	int *copy1, copy2[3];

	printf("copy1 = abc;\n");
	copy1 = abc;
	printf("copy1 now points to abc\n");
	printf("copying contents of abc to copy2\nmemmove(copy2, abc, sizeof(int)*3);\n");
	memmove(copy2, abc, sizeof(int)*3);

	printf("Changing abc[0] to 3\n");
	abc[0]=3;
	printf("assert(copy1[0]==3);\n");
	assert(copy1[0]==3);
	printf("assert(copy2[0]==0);\n");
	assert(copy2[0]==0);
};
