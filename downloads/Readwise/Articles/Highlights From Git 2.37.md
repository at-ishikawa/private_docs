# Highlights From Git 2.37

![](https://readwise-assets.s3.amazonaws.com/static/images/article4.6bc1851654a0.png)

### Metadata

- Author: Jeff Hostetler
- Full Title: Highlights From Git 2.37
- Category: #articles
- URL: https://github.blog/2022-06-27-highlights-from-git-2-37/

### Highlights

- In Git, we often talk about classifying objects as either “reachable” or “unreachable”. An object is “reachable” when there is at least one reference (a branch or a tag) from which you can start an object walk (traversing from commits to their parents, from trees into their sub-trees, and so on) and end up at your destination. Similarly, an object is “unreachable” when no such reference exists.
- But observant readers will notice the gc.pruneExpire configuration. This setting defines a “grace period” during which unreachable objects which are not yet old enough to be removed from the repository completely are left alone. This is done in order to mitigate a race condition where an unreachable object that is about to be deleted becomes reachable by some other process (like an incoming reference update or a push) before then being deleted, leaving the repository in a corrupt state
