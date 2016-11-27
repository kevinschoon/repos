## Repos

The majority of work I do on a day to day basis happens inside of a directory called `repos` in my `$HOME`.
The directory structure looks like this:

    ~/repos
      aur                            # Arch Linux AUR repositories
      go/src/github.com/kevinschoon  # GOPATH
      ks                             # Personal Repositories
      clients/*/**                   # Client/Work repositories
      ...


Repos is a simple tool to show dirty repositories across different directories.
Each top level directory is called a `Collection`.


### Installation

  go get github.com/kevinschoon/repos


### Configuration

  Top level repos can be configured inside a `~/.config/repos/config.json` file. Configuration is stored
  in JSON format and currently only has two options. Directories must be relative to `basePath`.

    {
      "basePath": "/home/kevin/repos",
      "collections": [
        {
          "name": "Go",
          "pattern": "go/src/github.com/*/*"
        },
        {
          "name": "Personal",
          "pattern": "ks/*"
        },
        {
          "name": "Clients",
          "pattern": "clients/*/*"
        }
      ]
    }

### Usage

Show the status of the repos directory

    $ repos

    > REPOSITORY                           PENDING  STASHED
    > go/src/github.com/kevinschoon/repos  3        0
    ...

List repos with pending changes

    repos -pending

    > fuu/my-repo
    > bar/my-other-repo

    wc -l $(repos pending)


List repos with stashed changes

  > repos -stashed

  > fuu/my-repo
  > bar/my-other-repo


### TODO

Better integration with Git
???

