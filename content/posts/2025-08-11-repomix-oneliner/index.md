---
title: "Repomix oneliner"
date: 2025-08-11T05:43:29-05:00
categories:
- tech
tags:
- tech-tip
---

Claude has been a pretty decent sidekick to pair program with these past few
months.

Using [Repomix](https://github.com/yamadashy/repomix) with Claude has definitely
made Claude more useful for software dev tasks.  Repomix bundles up all the code
in a given directory structure into an xml file.  I use this packaged code to
provide context to the LLM I'm working with.

Instead of manually baking repomix into each of my projects' flake
configurations, I created a fish alias that:
- uses nix shell to install repomix via nodejs npx into a temporary environment
- runs Repomix to stdout and copies the contents to the system clipboard

nix onliner:
```bash
nix-shell -p nodejs --run 'npx repomix --stdout --copy'
```

No doubt, there are limits to how useful this can be for larger codebases.
(primarily around security and utility)  I'm still also experimenting with
Claude Code and the like, but I prefer not to run these tools directly in my
primary development environment.  I sandbox them in NixOS generated VMs with the
code I want exposed.  This overhead translates to Repomix still having a place
in my toolbox for the time being.
