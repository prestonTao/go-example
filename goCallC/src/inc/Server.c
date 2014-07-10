/*
 ============================================================================
 Name        : Hello.c
 Author      : tao
 Version     :
 Copyright   : Your copyright notice
 Description : Hello World in C, Ansi-style
 ============================================================================
 */

#include <stdio.h>
#include <stdlib.h>

#include "foo.h"

// int main(void) {
// 	void socketConnet();
// 	void callLib();
// 	socketConnet();
// 	callLib();
// 	return EXIT_SUCCESS;
// }

void socketConnet(){
	float a,b;
	a = 123456.789e5;
	b = a + 20.0;
	printf("%f\n%f",a,b);
}

void callLib(){
	foo();
}
