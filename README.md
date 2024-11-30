
# setup

1. clone down the repo
```
git clone git@github.com:ruebenramirez/blog.git blog
cd blog
git submodule update --init
git lfs install
git lfs checkout
```

1. install nix
1. Install `direnv`
1. `echo "use flake" > .envrc`
1. `direnv allow`
  - this should build the flake


# Generating a new post

```shell
hugo new "posts/$(date --iso-8601)-post-title.md"
```




