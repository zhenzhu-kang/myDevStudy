## 24.06.22

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
기존 윈도우에 있던 리포지토리를 깃허브에 푸시하고 우분투에서 깃클론을 한뒤에 첫 push였는데 문제가 생겼었다..
`git config --global pull.rebase false`
```
`