#github #git

[[git 명령어]]
## 1. github에 파일 업로드 하기
(계정이 있다는 가정하에 진행)
1. github에서 새로운 레포지토리를 생성한다.
2. 터미널을 켜서 cd 명령어를 통해 github에 업로드하고 싶은 폴더로 이동한다.
3. 터미널에 `git init`을 입력한다. (초기화 작업, `git init`을 하지 않으면 git 명령어가 안 먹음)
4. 터미널에 `git add .` 를 입력, 그 다음 `git commit -m "하고싶은 코멘트 입력"` 를 입력한다.
5. 다시 github로 들어와서 새로 만든 레포지토리 페이지 맨 아래에 있는 명령어 3줄을 복사해서 터미널에 입력한다. 
	`git remote add origin https://github.com/zhenzhukang/devStudy.git`  
		> 내 레포지토리 페이지 주소를 origin이라는 명령어로 지정 (origin = 페이지 주소)
		> 이렇게 하면 일일이 페이지 주소를 치지 않고 origin만 쳐도 됨
	`git branch -M main`
		> 마스터 브렌치명을 main으로 지정
	`git push -u origin main`
		> `git push` 뒤에 `origin main` 안치고 `git push`만 쳐도 github에 push가 됨
6. 위 명령어 3줄을 입력하면 내가 업로드하고 싶은 폴더가 github에 push가 됨
7. 이제 push 하고 싶을 때 아래 명령어를 반복하면 됨.
	`git add .`
	`git commit -m "하고싶은 코멘트 입력"`
	`git push`
