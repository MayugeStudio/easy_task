# Easy-Task

The tool makes task visualization a breeze!

## Usage
For example, let's say you have a file named `example.md`

```markdown 
- [ ] Create project tracker
- [ ] Write Go code
- [X] Build visualization
- Learn React
    - [X] Go to react example site.
    - [ ] Create simple todo app.
```

In command line, simply use:

```shell
etask example.md
```

And voilà! It will display:

```text
[ ] Create project tracker
[ ] Write Go code         
[X] Build visualization
Learn React
  [X] Go to react example site.
  [ ] Create simple todo app.
  [##########          ]50%
[################                        ]40%
```

**🚀 Stay tuned for the next update!**
