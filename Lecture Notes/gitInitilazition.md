> If you want to create a repo from _scratch_ here is the milestones.

- Go to github.com or whatever your favourite git provider and create a repo. Do not add any **gitignore ...vs files.**
- In your locale open `Git bash` and head to your project directory.
- **`Git init`** -> This initializes your empty git repo
- From your favourite ide add  _.gitignore and README.md_ file.
- In Gitbash  write `git add -all`(_Add all new files_) or `git add .gitignore` (add specific file)
- **`git commit -m 'first commit'`**. Before committing you can see your git status via `git status`
- This command must be run just once. **`git remote add origin  https://github.com/username/reponame.git`**
This command ensures that your local repo connects to the right github repo.
- **`git push -u origin master`** _push your local repo to remote online repo_.

Thats all.
