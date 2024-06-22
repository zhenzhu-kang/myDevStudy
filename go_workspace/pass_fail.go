// pass_fail 성적 합격 여부 확인 프로그램
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin) //버퍼로부터 입력을 받는다.
	input, err /*에러 반환값을 err에 저장한다*/ := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) //err에 할당된 오류를 출력
	} //err변수값이 nil(정상)이 아닐시 에러를 보고
	fmt.Println(input)
}
