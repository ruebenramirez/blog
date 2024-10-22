---
title: "Display Configuration With Kanshi"
date: 2024-10-01T13:52:28-05:00
categories:
- tech
tags:
- wayland
- sway
- desktop
- linux
---

I switched from i3 (X11) to Sway (Wayland) on my NixOS machines.  There have been a number of issues I've worked through to make this happen, but today I'm documenting my monitor setup.

On i3 I controlled my monitors with some monitor configuration scripts that I generated with `arandr`.  `arandr` is a desktop gui app that provides a visual way to enable/disable available monitors, position monitors, configure resolution, and set a primary display, and save all of these as an xrandr bash script.  Using this workflow, I saved different monitor setup scripts for the different machines I use regularly.  It was brittle though and not as easy to update as I would have liked.

I was looking for a similar workflow as I switched over to Sway, and I found by default outputs are hardcoded into the ~/.config/sway/config file.  This is problematic if you have different monitors to plug into or if you share sway configuration across multiple machines (hardcoded values for the wrong displays is problematic).  I did find `wdisplays` which works similarly to `arandr`.  It offers a GUI for quick one-time configuration (e.g. if you're giving a presentation and just need a quick mirror configuration, this is probably the tool for you!)

`wdisplays` didn't help me with multiple monitor configuration profiles though. Example scenario for 3 monitor profiles: 1) a laptop that's used at coffee shops, 2) but also plugged in at a desk at home 3) or a desk in an office.  So I set out to find a better solution for dynamic monitor configuration and stumbled upon [`kanshi`](https://git.sr.ht/~emersion/kanshi).

Kanshi is a dynamic window configuration daemon that allows you to create many profiles in a configuration file, and allow it to determine which one to use based on which displays are connected to your machine.

##  setting up `kanshi`

I updated my nix config to capture kanshi display configuration service.
- warning: I'm not using home manager, so this is a little more involved (and I can see the appeal of offloading setup burdens like this to something like home manager).
- confirmed that sway is enabled in my nix flake

``` nix
programs.sway.enable = true;
```

- added `kanshi` to my nix flake system packages

``` nix
environment.systemPackages = with pkgs; [
    # ...
    kanshi
    # ...
];
```


### systemd unit to manage kanshi daemon runtime

Kanshi is just a daemon that needs some help managing runtime.
- created a user level systemd unit for the kanshi daemon in my nix flake
``` nix
  systemd.user.services.kanshi = {
    description = "kanshi dynamic display congfiguration daemon";
    wantedBy = [ "graphical-session.target" ];
    partOf = [ "graphical-session.target" ];
    after = [ "graphical-session.target" ];
    serviceConfig = {
      Type = "simple";
      ExecStart = ''${pkgs.kanshi}/bin/kanshi'';
    };
  };
```


### kicking off kanshi from sway config
I added this line into my `~/.config/sway/config` to restart kanshi any time the sway config loads.

``` bash
exec_always systemctl --user restart kanshi &
```

I have a keyboard shortcut to reload my sway config, which in turn will restart the kanshi unit.

``` bash
bindsym $mod+Shift+c reload
```

### creating kanshi display configuration
I came up with this kanshi config (`~/.config/kanshi/config` on my machine) to define profiles for running around with my laptop unplugged or when everything is plugged in at my desk:

``` c
profile thinkpad_undocked {
    output "California Institute of Technology 0x1404 Unknown" mode 1920x1200@60Hz position 0,0
}

profile thinkpad_x220_undocked {
    output "LG Display 0x036C Unknown" mode 1366x768 position 0,0
}

profile xps17_undocked {
    output eDP-1 mode 3840x2400@60Hz position 0,0
}

profile xps17_desk {
    output "Sharp Corporation 0x1517 Unknown" mode 3840x2400 position 3840,0 scale 2.00
    output "GWD ARZOPA " mode 2560x1600 position 3840,1200 scale 1.20
    output "LG Electronics LG HDR 4K 406NTZNA2149" mode 3840x2160@60Hz position 0,0 scale 1.00
}
```


#### using kanshi "criteria" or display name

Kanshi can use either a display name or what the docs call "criteria".  This is `make`, `model`, and `serial` as a single string, where `Unknown` replaces any missing/empty display values.

I had some problems getting kanshi to apply the correct profile when I was using "criteria" to target profile displays.  TIP: It turns out empty strings `""` like the `serial` field value for my usb-c monitor is not the same as `Unknown`.  Once I just left an extra space at the end of my `"GWD ARZOPA "` output criteria string, kanshi started playing nice with profile matching!

In order to identify display information from sway:
``` bash
$ swaymsg -t get_outputs
# outputs a LOT of information
```

to narrow down the display interface "name" or "criteria" of the displays I wanted to target:
``` bash
$ swaymsg -t get_outputs | egrep "name|make|model|serial"
    "name": "eDP-1",
    "make": "Sharp Corporation",
    "model": "0x1517",
    "serial": "Unknown",
    "name": "DP-2",
    "make": "GWD",
    "model": "ARZOPA",
    "serial": "",
    "name": "DP-4",
    "make": "LG Electronics",
    "model": "LG HDR 4K",
    "serial": "406NTZNA2149",
```

### How has this changed my workflow

On `i3` I used a `rescreen` shell script I wrote to apply the correct display configuration based on whichi monitors were available.  The actual monitor configurations were just more shell scripts generated by `arandr`, that my `rescreen` would call downstream.

Now on `sway` I use my sway configuration to restart the kanshi daemon.  Every time the kanshi daemon loads, it detects the connected displays and selects the appropriate kanshi dispaly profile from my profile configuration.  So it's now configuration instead of shell scripting that ties everything together.

In daily practice this means that I reload my sway configuration with a `mod + shift + c` keybinding which will kick off a kanshi daemon restart.  This is a lot simpler to me workflow than opening a terminal and running a shell script.  It's less distracting and creates a foundation for *"if there's something wrong, just reapply the sway config"*.

