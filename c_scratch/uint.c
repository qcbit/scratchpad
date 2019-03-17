#include <stdio.h>

int main(){
	int neg = -2;
	size_t zero = 0;
	if (neg < zero) printf("%d < %ld\n", neg, zero);
	else 		printf("%d > %ld\n", neg, zero);

	printf("most comparisons between signed and an unsigned int"\
		" C will force the signed type to unsigned\n");
}
