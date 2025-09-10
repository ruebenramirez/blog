---
title: "Python Dev in a Nix Flake"
date: 2025-08-27T15:20:50-05:00
categories:
- tech
tags:
- Nix
- Linux
- dev
---

My ideal development workflow: git clone a repo, cd into my local project dir,
and immediately run code/tests.

Normally this requires a LOT of effort: operating system dependencies,
programming language dependencies, libraries and frameworks and sometimes even
more.  For most of my career, development environments have been provided in the
form of virtual machines and then docker containers relying on imperative shell
script automation to setup part or even all of the environment.  This is usually
problematic and error prone though.

I've found Nix flakes and `direnv` incredibly helpful for managing my
development environments. Flakes provide reproducible environments through
declarative syntax while direnv makes it super easy to load the environment.


### A simple Nix flake

`flake.nix`
``` nix
{
  description = "Development shell with Python 3.12 and uv";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            python312
            uv
          ];
        };
      });
}
```

- `nixpkgs` packages from the Nix community
  - using unstable packages for development in this example
- `flake-utils` helper functions that keep the boilerplate to a
minimum while providing multiple CPU architecture support
- adding `python312` and `uv` Nix packages to the dev environment


`.direnv`
``` shell
use flake
watch_file flake.nix flake.lock
```
- Automates (re)loading the Nix flake when updated


`.gitignore`
``` shell
.direnv/**
```
- this will ensure the dependency files don't find their way into your repo


### The workflow

``` shell
cd <project directory>
```

``` shell
# only required on the first run
direnv allow
```
- enables the autoloading of the flake when entering this directory
- builds the flake
- generates a `flake.lock` which pins the python and uv dependencies
  - (update these pinned versions with the `nix flake update` command)

``` shell
# create a python `uv` project
uv init
```
- This creates uv python project files.

``` shell
# add a dependency
uv add requests
```
- installs the python dependency
- pins the dependency version in the `uv.lock` file


update our `main.py`
``` python
import requests

def main():
    r = requests.get('https://blog.rueb.dev/')
    print("response status code: {}".format(r.status_code))

if __name__ == "__main__":
    main()
```
- make use of the requests library we added to our project


Time to see our code in action!
```shell
uv run main.py
status code response: 200
```


### Why do any of this?

- Reproducibility: You'll have python and uv available, whether your an AMD
Linux geek or running on the fanciest Apple silicon money can buy.
- declarative syntax: define what the target environment should include and let
Nix "make it so"
- powerful ecosystem with #AllThePackages: I'll be very surprised the day I can't
find what I'm looking for at https://search.nixos.org


### Resources

- This is one of the best collections of Nix reference material I've come
across: https://tinkering.xyz/nix-docs/
- direnv docs: https://direnv.net
- Python uv package management: https://docs.astral.sh/uv/

### Next steps

I'm continuing to explore Nix capabilities with more simple examples:

- Migrating from uv to Nix managed python deps with `uv2nix`
- Packaging applications with Nix generated containers (Nix flake > Dockerfile)
- Interfacing with external services (databases, caching, object storage, etc.)

