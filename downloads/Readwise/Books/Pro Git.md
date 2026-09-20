# Pro Git

![](https://images-na.ssl-images-amazon.com/images/I/51N%2BrMcj75L._SL200_.jpg)

### Metadata

- Author: Scott Chacon and Ben Straub
- Full Title: Pro Git
- Category: #books

### Highlights

- If you don’t want to type it every single time you push, you can set up a “credential cache.” The simplest is just to keep it in memory for a few minutes, which you can easily set up by running git config --global credential.helper cache. ([Location 1969](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=1969))
- rebasing makes for a cleaner history. If you examine the log of a rebased branch, it looks like a linear history: it appears that all the work happened in series, even when it originally happened in parallel. ([Location 2084](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2084))
- git rebase --onto master server client ([Location 2103](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2103))
- Do not rebase commits that exist outside your repository. ([Location 2128](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2128))
- In general the way to get the best of both worlds is to rebase local changes you’ve made but haven’t shared yet before you push them in order to clean up your story, but never rebase anything you’ve pushed somewhere. ([Location 2190](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2190))
- git instaweb --httpd=webrick ([Location 2563](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2563))
- Instead of “I added tests for” or “Adding tests for,” use “Add tests for.” ([Location 2787](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=2787))
- Note  You may want to use rebase -i to squash your work down to a single commit, or rearrange the work in the commits to make the patch easier for the maintainer to review. ([Location 3015](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=3015))
- git diff master...contrib ([Location 3316](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=3316))
- The Git project has four long-running branches: master, next, and pu (proposed updates) for new work, and maint for maintenance backports. ([Location 3349](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=3349))
- Tagging Your Releases ([Location 3401](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=3401))
- you want to make the trunk subdirectory be the new project root for every commit, filter-branch can help you do that, too: $ git filter-branch --subdirectory-filter trunk HEAD ([Location 5122](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5122))
- git filter-branch --commit-filter '         if [ "$GIT_AUTHOR_EMAIL" = "schacon@localhost"];         then                 GIT_AUTHOR_NAME="Scott Chacon";                 GIT_AUTHOR_EMAIL="schacon@example.com";                 git commit-tree "$@";         else                 git commit-tree "$@";         fi' HEAD ([Location 5134](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5134))
- Actually, it’s a bit smarter than that—it tries to do a trivial merge in the Working Directory, so all the files you haven’t changed in will be updated. reset --hard, and on the other hand, will simply replace everything across the board without checking. ([Location 5315](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5315))
- you see that you have a lot of whitespace issues in a merge, you can simply abort it and do it again, this time with -Xignore-all-space or -Xignore-space-change. The first option ignores changes in any amount of existing whitespace, the second ignores all whitespace changes altogether. ([Location 5408](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5408))
- You can pass --conflict either diff3 or merge (which is the default). ([Location 5516](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5516))
- The -m 1 flag indicates which parent is the “mainline” and should be kept. ([Location 5627](https://readwise.io/to_kindle?action=open&asin=B00LPDVAX2&location=5627))
