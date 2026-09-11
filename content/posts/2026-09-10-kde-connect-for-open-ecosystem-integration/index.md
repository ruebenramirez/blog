---
title: "KDE Connect for Open Ecosystem Integration"
date: 2026-09-10T15:32:16-05:00
categories:
- tech
tags:
- tech-tip
---

![KDE Connect showing SMS messages on desktop](./screenshot-2026-09-10-20:02:05.png)

I recently degoogled my phone. There's a lot of noise around the anti big-tech
scene these days, but I'm not as driven by that. I just prefer open source where
I can adopt it in my life. I'm the kid that grew up getting in trouble for
taking the TV remote control apart, and graduated to doing the same to our
family PC. So anyways, I uninstalled Google Play Services, disabled RCS chat,
removed Google Messages, and went back to standard SMS and MMS.

One thing I immediately missed was the ability to txt msg from my laptop.

I've never really loved doing things on my phone.  I started out in the Apple
ecosystem, using a MacBook to send messages through my connected iPhone. When I
switched away from macOS, I swapped from an iPhone to Android and could send
texts through Google Messages (so I ended up trading Apple for Google). It was
*more* open, since I could use a browser on my Linux laptop to send messages
through my phone. That felt like a step in the right direction.

And now I'm at it again. I swapped Android for GrapheneOS and cut out the few
remaining Google services I depended on. Then I learned about KDE Connect.
I'm a big fan of the project's mission and I'm impressed with how well it
functions.

With KDE Connect, I get to continue down this road of staying in touch without
living on my phone all day. There are no proprietary services. No vendor lock
in. Just open tools that work.

## The Setup

I installed KDE Connect from F-Droid on my GrapheneOS phone and added it to my
desktop. Since I manage my system with a NixOS flake, the desktop install was
a one line change:

```nix
programs.kdeconnect.enable = true;
```

([system flake commit
here](https://github.com/ruebenramirez/systems/commit/572bcdac67d8ad72cabeeeccd545db672043539c))

Pairing was seamless.

## A Minor Bug and How I Worked Around It

Open source software gets a bad rap for "when you are on the cutting edge, you
are going to bleed." I'm lucky to have AI tools that are much better at
searching through the dark alleys of the internet for fixes.

SMS messages weren't showing up on my desktop. Turns out I'm not alone. A
community [bug report](https://bugs.kde.org/show_bug.cgi?id=523424) went up just
a week before mine. A few adb commands later and I was all patched up. I'm
really impressed with how much functionality this thing has. Way more than I was
even looking for.

## Connecting From Different Networks

KDE Connect works great over LAN at home and over WiFi hotspot when I'm working
remotely. Device discovery happens automatically via mDNS. Once paired,
reconnecting across networks is seamless.

I'm also looking into forcing my devices to talk solely over a WireGuard subnet.
My personal backlog keeps growing.

## Features I Have Not Tried Yet

KDE Connect does way more than SMS relay. Clipboard sharing, remote input with
phone as trackpad, file transfer, media control, notification mirroring, even my
phone's battery level on the desktop. I have not explored all of these yet, but
It's awesome that all of this is available.
