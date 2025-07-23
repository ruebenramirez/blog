
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
hugo new "posts/$(date --iso-8601)-post-title/index.md"
```

I added this fish shell function to my `~/.config/fish/config.fish`
```
function hugopost
    set -l title (string replace -a ' ' '-' $argv[1])
    hugo new "posts/$(date --iso-8601)-$title/index.md"
end
```

this allows me to
```
hugopost "Martian Marshmallows"
```

which will create a new directory (instead of a single markdown file) and I can
then load all of my post assets (into this directory easily)




