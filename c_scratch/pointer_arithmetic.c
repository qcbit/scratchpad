#include <stdio.h>

int main(){
	char *list[] = {"first", "second", "third", NULL};
	for (char **p = list; *p != NULL; p++){
		printf("%s\n", p[0]);
	}

	// Implementing without knowing p++
	
	for (char **p = list; *p != NULL; p=p+1){
		printf("%s\n", p[0]);
	}
};
