# Configure Git 
https://git-scm.com/book/de/v1/Los-geht%E2%80%99s-Git-konfigurieren

``` 
git config --global user.name "John Doe"
git config --global user.email johndoe@example.com
```

### Store Credentials permanent in Linux / macOS us

```
git config --global credential.helper cache
```

### Clone repro localy from remote 

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

```
mkdir newDirectory
cd newDirectory
git init .

git status
```
On branch master

### Update a file / commit / push to GitHub


```
notepad readme.md

git add .

git status

git commit -m "Update git readme.md"

git pull

### Commit without extra add files 

git commit -a -m "commit without extra git add . command"
```

### Show git history

```
git reflog

git log

git log --graph
```

## diff changes 

```
git log --oneline --decorate
git diff <no> HEAD
```

```
git blame <file name>
```

## Branches

```
git branch
* master

git branch -av
(shows remove branches)
```

## Tags


```
git ls-remote --tags origin

```

## Git and WSL

To set up Git Credential Manager for use with a WSL distribution, open your distribution and enter this command:
https://docs.microsoft.com/en-us/windows/wsl/faq
```
git config --global credential.helper "/mnt/c/Program\ Files/Git/mingw64/libexec/git-core/git-credential-manager.exe"
```
