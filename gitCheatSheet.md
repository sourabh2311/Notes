### Creating a repo
* setup ssh on ur machine, follow this link: https://help.github.com/en/articles/connecting-to-github-with-ssh or basically follow this: https://www.youtube.com/watch?v=HfTXHrWMGVY
* Also the following is helpful in case working with multiple github accounts: https://medium.com/@xiaolishen/use-multiple-ssh-keys-for-different-github-accounts-on-the-same-computer-7d7103ca8693
* Two methods to create a repo:
  * Method 1:-
    * to create a repo on github and then clone it using ssh
  * Method 2:- To create a repo from an existing project and then connect it to github
    * Go into the directory containing the project. 
    * Type `git init`.
    * Type `git add .`.
    * Type `git commit`
    * Go to github.
    * Click the new repository button in the top-right. You’ll have an option there to initialize the repository with a README file, but I don’t.
    * Click the “Create repository” button.
    * `git remote add origin git@github.com:username/new_repo.git`
    * `git push -u origin master`
* `git checkout -b sourabh/neoLog/zap` - To create a new branch and then checkout to it
* `git add .` - To have all your files staged
* `git commit -m "msg"` - To commit
* `git push` - To push
* `git pull` - To pull
* `git branch` - To check what is your branch
* `git status` - Status
* `git diff` - See your changes
* `git stash` - To put changes onto a stack
* `git log` - See the log
* `git stash pop` - Put the changes in stack to the current branch
* `git stash show` - Show the changes in stack
* `git rebase eng` - https://www.youtube.com/watch?v=CRlGDDprdOQ  Basically in your master first do git pull, then change branch to your feature implementation branch and then type git rebase eng so that the base of your feature is taken as the latest commit in eng also you can then checkout to the eng branch and then type git rebase yourFeatureBranch to actually merge your changes.