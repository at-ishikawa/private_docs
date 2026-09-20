# Medium: [Print Folder Structure](https://interviewing.io/questions/print-folder-structure)

## Print Folder Structure Problem

Given a list of file paths, print all of the files in each of the folders.

For example:

Input: files = [ "/webapp/assets/html/a.html", "/webapp/assets/html/b.html", "/webapp/assets/js/c.js", "/webapp/index.html" ]

Output:
```
-- webapp
  -- assets
    -- html
      -- a.html
      -- b.html
    -- js
      -- c.js
  -- index.html
```

### 2024/8/24

- Communications: thinking a loud is good
    - Checking requirements with an interviewer more
        - How to recognize a file?
        - If an input can be a directory or not?
        - Input/output/function specs
- BFS: FIFO (queue, not stack)
