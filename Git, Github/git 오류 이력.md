OS: 우분투

```git title:24.06.22.error_분석
//git push 중 생긴 오류
To https://github.com/zhenzhu-kang/myDevStudy.git ! [rejected] main -> main (fetch first) 
error: failed to push some refs to 'https://github.com/zhenzhu-kang/myDevStudy.git'

hint: Updates were rejected because the remote contains work that you do 
hint: not have locally. This is usually caused by another repository pushing 
hint: to the same ref. You may want to first integrate the remote changes 
hint: (e.g., 'git pull ...') before pushing again. 
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

// 그래서 이번엔 pull을 해봤는데 아래 같은 오류가 떴다...
git pull origin main From https://github.com/zhenzhu-kang/myDevStudy * branch main -> FETCH_HEAD 

- hint: You have divergent branches and need to specify how to reconcile them. 
- hint: You can do so by running one of the following commands sometime before 
- hint: your next pull: 
- hint: hint: git config pull.rebase false # merge (the default strategy) 
- hint: git config pull.rebase true # rebase 
- hint: git config pull.ff only # fast-forward only 
- hint: hint: You can replace "git config" with "git config --global" to set a default 
- hint: preference for all repositories. You can also pass --rebase, --no-rebase, hint: or --ff-only on the command line to override the configured default per
- hint: invocation. fatal: Need to specify how to reconcile divergent branches.

//chat gpt의 추천으로 아래 명령어를 쳤더니 push가 성공했다.
`git config --global pull.rebase false`

기존 윈도우에 있던 리포지토리를 깃허브에 푸시하고 우분투에서 깃클론을 한뒤에 첫 push였는데 문제가 생겼었다..

/*이 오류는 로컬 브랜치와 원격 브랜치가 서로 다르게 진행된 경우에 발생하며, Git이 어떻게 이 두 브랜치를 병합해야 할지 명확하지 않기 때문에 발생합니다. 이 문제를 해결하려면 병합 전략을 지정해 주어야 합니다. 다음 단계들을 따라 문제를 해결할 수 있습니다. (chat gpt)*/

아마 윈도우에서 옵시디언으로 자동 커밋 푸시를 하고 있었는데 그거랑 우분투에서 푸시하는게 겹쳐서 그런걸까..?

/*`git config --global pull.rebase false` 명령어는 전역(Git 전체) 설정에서 `pull.rebase` 옵션을 설정하는 것입니다. 이 명령어를 사용하면 모든 Git 저장소에서 `git pull` 명령어를 실행할 때 기본적으로 리베이스를 사용하지 않고 기본 병합 전략을 사용하도록 설정됩니다.

여기서 각 부분의 의미는 다음과 같습니다:

- `git config --global`: Git의 전역 설정을 변경하겠다는 의미입니다. 즉, 이 설정은 사용자의 컴퓨터에서 모든 Git 저장소에 적용됩니다.
    
- `pull.rebase false`: `pull` 명령어 실행 시 리베이스를 사용하지 않도록 설정합니다. 따라서 기본적으로 병합(merge) 전략을 사용하게 됩니다.
    

이 설정을 적용하면 개별 Git 저장소에서 별도로 `git pull` 명령을 실행할 때 `--no-rebase` 옵션을 지정하지 않아도 리베이스를 사용하지 않고 병합을 수행하게 됩니다. (chat gpt)*/

pull.rebase false 는 곧 pull 기본설정을 rebase로 하지 않고(false) 병합으로 하겠다는 뜻..?

```
`