---
title: "Keeping Downloads Safe With Checksums"
date: 2024-06-20T14:38:41-05:00
categories:
- tech
tags:
- tech-tip
---

When we download a file from a trusted source, it's become standard practice for websites to provide a file checksum.  That checksum can help us ensure the file we download is the trusted version from them and that it has not been manipulated or infected with malware.  These file checksums are a fingerprint of a file.  We can use that fingerprint to make sure we're downloading the trusted, safe, file.

Said more simply: compare the fingerprint of the file we downloaded to the fingerprint that our trusted friends said we should have.


### Example case

VLC media player is one of my favorite ways to watch videos or play audio on my computers.  The VLC team offers the windows 64 installer download on their website here: https://get.videolan.org/vlc/3.0.21/win64/vlc-3.0.21-win64.exe

On this same page, they provide a checksum that allows you to confirm that the file you've downloaded is the trusted file they released.  If the checksums don't match, then we might assume someone might have inserted malware or done something dangerous to the file we've downloaded.  Confirming the download's checksum keeps everyone VLC users safe.

<a href="https://get.videolan.org/vlc/3.0.21/win64/vlc-3.0.21-win64.exe"><img src="/images/2024-06-20-vlc-file-download-checksum-example.png" alt="VLC file download checksum example" width="800" /></a>


### How do we do this though?

I provide a few examples on how to checksum files on different Operating Systems below.  There are plenty of other options available, including GUI tools.  The important thing is to build the habit of comparing the checksum of the files your download though!


#### on Windows

We use the PowerShell terminal to generate a SHA256 checksum of the file we downloaded.

Change directories to wherever you downloaded the file and then run PowerShell's handy little `Get-FileHash` command:

```powershell
PS D:\Downloads> Get-FileHash .\vlc-3.0.21-win64.exe -Algorithm SHA256

Algorithm       Hash                                                                   Path
---------       ----                                                                   ----
SHA256          9742689A50E96DDC04D80CEFF046B28DA2BEEFD617BE18166F8C5E715EC60C59       D:\Downloads\vlc-3.0.21-win64.exe
```


#### on Linux

We use a standard terminal environment to run the `sha256sum` command on the file we downloaded:

```shell
[/mnt/d/Downloads]$ sha256sum ./vlc-3.0.21-win64.exe
9742689a50e96ddc04d80ceff046b28da2beefd617be18166f8c5e715ec60c59  ./vlc-3.0.21-win64.exe
```


#### on MacOS

We use a standard terminal environment to run the `shasum` command on the file we downloaded:

```shell
shasum -a 256 ./vlc-3.0.21-win64.exe
9742689a50e96ddc04d80ceff046b28da2beefd617be18166f8c5e715ec60c59  ./vlc-3.0.21-win64.exe
```

