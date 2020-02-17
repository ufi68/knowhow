```

git config --global user.email "e@mail"
git config --global user.name "user"

mkdir ufi68
cdufi68
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
