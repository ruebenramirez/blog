---
title: "Generating Mermaids"
date: 2026-09-11T10:45:17-05:00
categories:
- tech
tags:
- tech-tip
---

It's difficult balancing the importance of maintaining documentation with the
pain required to write it.  Or at least it was before capable LLMs became so
accessible.

I originally learned about mermaid diagrams a couple of years ago.  I
immediately thought, "What a great idea to capture diagrams with code to allow
version controlling the docs in the same repo as the software itself". Then I
closed the website and never got around to actually writing mermaid code.

GenAI makes this a no brainer now, so I decided to give it another shot: I
spun up an opencode agent session and had it generate a mermaid diagram of some
home network changes I'm tinkering with.


### A simple Mermaid diagram of some machines on my home network

![Diagram of machines and services on my home network](./home-network.png)

mermaid diagram code:
```mermaid
graph TD
    A[AT&T Router] --- B[Internet]

    subgraph Main Network
        A --- C[Laptops]
        A --- D[Tablets]
        A --- E[Phones]
        A --- F[GL.iNet Router]
    end

    subgraph Subnetwork
        F -.WiFi.- R[Roku]
        F -.Ethernet.- G[XPS Machine]
        G --- H[DAS Storage\nZFS ZPool]
        G --- I[Jellyfin Service]
        G --- J[Samba file shares]
    end

    R -.- I

    C -.- K[Tailscale]
    D -.- K
    E -.- K
    K -.- I
    K -.- J

    classDef devices fill:#d4f1f9,stroke:#05445E,stroke-width:2px,color:#05445E;
    classDef services fill:#ffecb3,stroke:#663d00,stroke-width:2px,color:#663d00;
    classDef network fill:#d8f3dc,stroke:#1b4332,stroke-width:2px,color:#1b4332;

    class A,F,C,D,E,G,H,R devices;
    class I,J,K services;
    class B network;
```

I wouldn't bother to come up with this on my own, but with an LLM doing the
heavy lifting, it's easy to improve project docs with diagrams as code.
