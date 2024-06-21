[[C 이론 공부]]

```c title:stdio.h헤더함수
#include <stdio.h>

// 문자 입출력 함수
getchar() //하나의 문자를 받아서 반환한다. 버퍼를 사용한다
putchar() //하나의 문자를 받아서 출력한다.
getch() //하나의 문자를 받아서 반환한다. 버퍼를 사용하지 않는다.
putch() //하나의 문자를 받아서 출력한다.

scanf() //
printf() //
gets(char s[], int length) //한 줄의 문자열을 읽어서 배열 s[]에 저장한다.
puts(char s[]) //배열 s[]에 저장되어 있는 한줄의 문자열을 출력한다.

#include <stdio.h>
#define SIZE 15
int main(void)
{
	char string[SIZE];
	gets(string, SIZE); //string이라는 변수에 문자를 입력받는다. SIZE 길이 만큼의 문자열을 입력받는다.
	puts(string)
}

```
```c title:string.h헤더함수
#include <string.h>

strlen(s) //배열의 크기가 아닌 실제 문자열 s의 길이를 반환한다
strcpy(s1, s2) //s1애 s2를 복사한다.
strcat(s1, s2) //s1의 끝에 s2를 붙여넣는다.
strcmp(s1, s2) //s1과 s2를 비교한다. 비교 일치하면 0을 반환
atoi(s) //문자열 s를 int형으로 반환한다.

```

```c title:window.h헤더함수
Sleep(1000);//1초 쉼 
system("cls"); //화면지움

```
```c title:math.h헤더함수
#include <math.h>
sin() //sin값 계산
cos() //cos값 계산
tan() //tan값 계산

int main()
{
	print("sin = %.2lf\n", sin(2))

	double data = 3.2;
	printf("cos = %.2lf\n", cos(data));
	double result = tan(4)+5;
	printf("tan = %.2lf\n", result);
}
```

```c title:stdlib.h헤더함수
#include <stdlib.h>
#include <time.h>

srand(time(NULL)); //time(NULL)은 현재 시간을 반환함. srand를 통해 난수의 시드를 설정
rand() //1~50까지 랜덤 숫자 반환

```