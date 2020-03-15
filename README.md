# Repos

A tiny script to recursively list git repositories under a given path.

## Usage

```
export REPOS_PATH="$HOME/repos"
alias r="cd \$(repos | fzf)"
r
```
