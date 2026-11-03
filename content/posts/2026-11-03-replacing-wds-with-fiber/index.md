---
title: "Replacing WDS with Fiber"
date: 2026-11-03T15:40:51-06:00
categories:
- tech
---

TL;DR: I did a thing with fiber and surprise surprise, it's better than
wireless.

In my last post I was really excited about buying a second router, setting up
WDS, and moving some server gear to a better spot in my apartment.  I had
considered it a win that I didn't have to mess with ethernet cabling.  Well that
didn't last long.

I've now replaced the WDS-based network setup with a dedicated fiber run.  It
started as an effort to address intermittent network issues, but ended up
improving nearly every aspect of my home network, from transfer speeds and Wi-Fi
performance to overall responsiveness.

### The Problem with WDS

My previous network configuration relied on Wireless Distribution System (WDS)
to bridge parts of my network.  It was great under normal conditions, but
dropped off during heavy usage.  Large data transfers would cause servers
sitting behind the WDS connection to become unresponsive.

I'm working with LLMs on my Strix Halo machine.  I'm experimenting with several
models (Qwen 3.6 35B, Qwen 3.8 27B, and Qwen 3.8 Flash Next) on different
inference engines (llama.cpp, Halogen, and Gufo).  These models are relatively
large (anywhere from 14-110GB), and moving model data around the network
generates substantial traffic.

During these transfers, the network would become unresponsive.  Even my Mosh
sessions to servers behind the WDS connection would disconnect, sometimes for as
long as 12 seconds.  It was pretty clear that the wireless bridge was a
reliability problem.

It's totally possible my network problems are due to the specific hardware I'm
running (2 * Flint 2 M6000 wifi routers) or bugs in the OpenWRT firmware (I did
at least confirm the latest firmware was installed, though), or that I'm in an
apartment where there are just a TON of other wireless networks setup and
potentially conflicting with my setup.

### Moving to Fiber

Rather than continuing to troubleshoot the wireless bridge, I decided to
eliminate it from the network's backhaul entirely.

I picked up some more hardware:
- Another "dumb" network switch
- SFP+ fiber optic transceivers
- some fiber optic cable

For surprisingly few bucks, I was able to replace the WDS connection with a
dedicated fiber run between the "dumy" switches.

One nice advantage of fiber is its physical size.  Compared to traditional
Ethernet cabling, fiber is considerably thinner and easier to route and hide.  I
was able to run the cable out of the way and keep the installation relatively
unobtrusive.  This was my first time working with fiber, and it was far easier
to run and setup than I anticipated.

### Giving Wi-Fi Its Own Job

Now that the WDS link was gone, I had a spare Wi-Fi router to put to work. I set
it up as a dedicated access point in the living room, where most of our devices
are used, and disabled Wi-Fi on the main router to reduce overhead.  (The theory
here being that a cooler router is a higher performing router.)

Because the fiber line handles the backhaul, the access point only has to worry
about serving Wi-Fi.

### The Results

![mind blown](./mindblown.gif)

The improvements have been noticeable across the board:

- **Wi-Fi performance improved significantly.**  Simply relocating the access point to
the living room resulted in much faster wireless speeds.
- **Data transfer speeds nearly doubled.**  Separating the wireless access point
from the network backhaul has made a substantial difference in throughput.
Sure, I expected an obvious performance improvement from replacing wireless
connectivity with dedicated fiber.  But doubling speeds was very unexpected and
very cool.
- **The main router runs cooler.**  Removing the Wi-Fi access point
responsibilities from the main router resulted in noticeably lower operating
temps.
- **The network is considerably more responsive.**  The network just feels
snappier and connections to servers are, once again, reliable.  My Mosh sessions
remain connected and servers on the backhaul run are no longer intermittently
disappearing while large volumes of data are on the move.

### Final Thoughts

This was one of those projects where the benefits extended well beyond the
original problem.  By separating routing, wireless access, and the network
backhaul into more distinct responsibilities, the setup is faster and far more
reliable now.

I wish I had just set out to run the fiber in the first place.  This has
absolutely been worth the effort to revisit though.
