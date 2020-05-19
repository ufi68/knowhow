# Configure Git 
https://git-scm.com/book/de/v1/Los-geht%E2%80%99s-Git-konfigurieren

``` 
git config --global user.name "John Doe"
git config --global user.email johndoe@example.com
```

```
mkdir ufi68
cd ufi68
git clone https://github.com/ufi68/knowhow.git
cd knowhow
git status

git remote  --verbose

git branch --list --remote --verbose

git fetch

git branch new_branch

git checkout new_branch

git commit -a -m "Add a commit message"

git push --set-upstream new_branch
git push -u origin new_branch

git log --oneline --decorate
```

### initilaize a new repro local

mkdir newDirectory
cd newDirectory
git init .

git status
On branch master

### Update a file / commit / push to GitHub


```
notepad readme.md

git add .

git status

git commit -m "Update git readme.md"

git pull
```









