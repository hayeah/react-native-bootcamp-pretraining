# Git Knowledge Checklist

We are going use Git for version control. Even if you are just working by yourself, you should still use Git. The reasons are:

+ It helps you remember what you've done.
+ You don't have to worry about breaking things by making the wrong changes. Easy to rollback.
+ It's the easiest way to publish your code for others to see.
+ Keeps clean by deleting code that you don't use. Since they are in the history, you don't "lose" the code.

If you don't have a GitHub account yet, go [sign up for GitHub](https://github.com/join).

## Git Basics

First, you'll need to know how to create or download an existing repository.

+ Can you use `git clone` to download a repository from GitHub?
+ Can you use `git init` to create a new repository?
+ What is the `.git` directory?

[创建版本库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743256916071d599b3aed534aaab22a0db6c4e07fd0000)

Then you make changes to the repository by adding new files or editing old files.

+ Can you use `git diff` and `git status` to see what changes you've made?
+ What do `git add` and `git commit` do?
+ Can you use `checkout` to go to a commit?
+ Can you use `git rm` to remove a file from source control?
+ If you've broken something, do you know how to use `git reset` to rollback?

[时光机穿梭](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743858312764dca7ad6d0754f76aa562e3789478044000)

## Remote Repositories

You can publish your code on the Internet so people can see your code, and make changes to your code.

+ What is a remote branch? What's the difference between `master` and `origin/master`?
+ Can you add a remote repository using `git remote add`?
+ Can you use `git push` to publish your code to GitHub?
+ Can you use `git pull` to get new commits from a remote branch?

[远程仓库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001374385852170d9c7adf13c30429b9660d0eb689dd43a000)

## Git Branching

Although you can ALWAYS work on the `master` branch, it's often convenient to start new branches for each feature you do. This is especially helpful if you work in a team.

+ Can you use `git branch` to create a new branch?
+ Can you use `git checkout` to switch between branches?

[Feature分支](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001376026233004c47f22a16d1f4fa289ce45f14bbc8f11000)

## 客户端

Once you know the basics, you could try one of the Git desktop clients.

+ [SourceTree](https://www.sourcetreeapp.com/)
+ [Git Tower](http://www.git-tower.com/)

For more complicated stuff, you'd need to go back to the command line.