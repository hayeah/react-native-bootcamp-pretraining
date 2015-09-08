# Git Knowledge Checklist

We are going use Git for version control. Even if you are just working by yourself, you should still use Git. The reasons are:

+ It helps you remember what you've done.
+ You don't have to worry about breaking things by making the wrong changes. Easy to rollback.
+ It's the easiest way to publish your code for others to see.
+ Keeps clean by deleting code that you don't use. Since they are in the history, you don't "lose" the code.

If you don't have a GitHub account yet, go [sign up for GitHub](https://github.com/join).

<cn>
# Git 知识清单

我们会使用 Git 来进行版本控制。即使只是自己工作，你也应该使用 Git。原因：  

+ 它帮助你记住你做了什么。
+ 你不需要担心因为错误的改动破坏了代码。回滚是非常简单的。
+ 这是对别人发布你的代码的最简单方式。
+ 通过删除你不用的代码来保持整洁。因为它们在修改历史中，你不会“丢失”这些代码。

如果你还没有一个 GitHub 账号，前往 [注册 GitHub](https://github.com/join)。
</cn>

## Git Basics

First, you'll need to know how to create or download an existing repository.

+ Can you use `git clone` to download a repository from GitHub?
+ Can you use `git init` to create a new repository?
+ What is the `.git` directory?

[创建版本库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743256916071d599b3aed534aaab22a0db6c4e07fd0000)

<cn>
## Git 基础

首先，你需要知道如何去创建或者下载一个现有的仓库。

+ 你能用 `git clone` 下载 GitHub 上的一个仓库吗？
+ 你能用 `git init` 创建一个新的仓库吗？
+ `.git` 目录是什么？

[创建版本库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743256916071d599b3aed534aaab22a0db6c4e07fd0000)
</cn>

Then you make changes to the repository by adding new files or editing old files.

+ Can you use `git diff` and `git status` to see what changes you've made?
+ What do `git add` and `git commit` do?
+ Can you use `checkout` to go to a commit?
+ Can you use `git rm` to remove a file from source control?
+ If you've broken something, do you know how to use `git reset` to rollback?

[时光机穿梭](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743858312764dca7ad6d0754f76aa562e3789478044000)

<cn>
然后通过添加新文件或者修改旧文件来做出改动。

+ 你能使用 `git diff` 和 `git status` 查看你做过什么修改么？
+ `git add` 和 `git commit` 是做什么的？
+ 你能使用 `checkout` 来切换到一次提交吗？
+ 你能使用 `git rm` 来从源代码控制中移除一个文件吗？

[时光机穿梭](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/0013743858312764dca7ad6d0754f76aa562e3789478044000)
</cn>

## Remote Repositories

You can publish your code on the Internet so people can see your code, and make changes to your code.

+ What is a remote branch? What's the difference between `master` and `origin/master`?
+ Can you add a remote repository using `git remote add`?
+ Can you use `git push` to publish your code to GitHub?
+ Can you use `git pull` to get new commits from a remote branch?

[远程仓库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001374385852170d9c7adf13c30429b9660d0eb689dd43a000)

<cn>
## 远程仓库

你可以把你的代码发布到网络上，让人们可以看到你的代码，对你的代码做出改动。

+ 远程分支是什么？`master` 和 `origin/master` 之间的区别是什么？
+ 你能用 `git remote add` 添加一个远程仓库吗？
+ 你能用 `git push` 发布你的代码到 GitHub 吗？
+ 你能用 `git pull` 从远程分支获取新的提交吗？

[远程仓库](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001374385852170d9c7adf13c30429b9660d0eb689dd43a000)
</cn>

## Git Branching

Although you can ALWAYS work on the `master` branch, it's often convenient to start new branches for each feature you do. This is especially helpful if you work in a team.

+ Can you use `git branch` to create a new branch?
+ Can you use `git checkout` to switch between branches?

[Feature分支](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001376026233004c47f22a16d1f4fa289ce45f14bbc8f11000)

<cn>
## Git 分支
尽管你可以**经常**在 `master` 分支工作，为每个你做的每个特性添加新的分支经常是非常方便的。如果你在一个团队中工作，这是非常有用的。

+ 你能使用 `git branch` 创建一个新的分支吗？
+ 你能使用 `git checkout` 在分支间切换吗？

[Feature分支](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000/001376026233004c47f22a16d1f4fa289ce45f14bbc8f11000)
</cn>

## 客户端

Once you know the basics, you could try one of the Git desktop clients.

+ [SourceTree](https://www.sourcetreeapp.com/)
+ [Git Tower](http://www.git-tower.com/)

For more complicated stuff, you'd need to go back to the command line.

<cn>
## 客户端

只要你了解了这些基础知识，你可以尝试这些 Git 桌面客户端中的一个。

+ [SourceTree](https://www.sourcetreeapp.com/)
+ [Git Tower](http://www.git-tower.com/)

对于更复杂的东西，你需要回到命令行。
</cn>