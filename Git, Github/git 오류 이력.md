## 24.06.22


```git title:24.06.22.error_분석
//git push 중 생긴 오류
To https://github.com/zhenzhu-kang/myDevStudy.git ! [rejected] main -> main (fetch first) 
error: failed to push some refs to 'https://github.com/zhenzhu-kang/myDevStudy.git'

hint: Updates were rejected because the remote contains work that you do 
hint: not have locally. This is usually caused by another repository pushing 
hint: to the same ref. You may want to first integrate the remote changes 
hint: (e.g., 'git pull ...') before pushing again. 
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

zhenzhu_kang@DESKTOP-109HCKC:~/myDevStudy$ git pull origin main From https://github.com/zhenzhu-kang/myDevStudy * branch main -> FETCH_HEAD 
- hint: You have divergent branches and need to specify how to reconcile them. 
- hint: You can do so by running one of the following commands sometime before 
- hint: your next pull: 
- hint: hint: git config pull.rebase false # merge (the default strategy) 
- hint: git config pull.rebase true # rebase 
- hint: git config pull.ff only # fast-forward only 
- hint: hint: You can replace "git config" with "git config --global" to set a default 
- hint: preference for all repositories. You can also pass --rebase, --no-rebase, hint: or --ff-only on the command line to override the configured default per
- hint: invocation. fatal: Need to specify how to reconcile divergent branches.
```
`