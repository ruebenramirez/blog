Title: light: a backlight controller 
Date: 02-09-2015 17:12
Category: sys config 

I have a 2012 15" unibody macbook pro with both intel and nvidia display controllers.  I get better performance out of the nvidia controller so I made it my default since I typically use this machine as a desktop at the house.

I ran into a problem while trying to dim the backlight today though.  On my other laptops, MacBook Pro Retina (from work) included, I use xbacklight via keyboard shortcuts...and it's worked flawlessly.

As it turns out, my older Macbook Pro has a dedicated acpi backlight controller that isn't controlled with my xorg configuration.  source: https://wiki.ubuntu.com/Kernel/AppleGmuxBacklight

I did a lot of digging and found that most of the solutions online offer some variant of, go into your xorg.conf files (and if they don't exist, create them) and then pray to the spaghetti monster that it works.

I found a nice alternative called [light](https://github.com/haikarainen/light/) on the [ArchWiki](https://wiki.archlinux.org/index.php/backlight#light).  The really nice thing about `light` is that there's no dependency on your xorg configuration.  It works independently to talk to your hardware to dim or brighten up your display.

I was able to install and remap my shortcuts in just a few line edits in my dotfiles: [https://github.com/ruebenramirez/.dotfiles/commit/a18fd396a4f53238c7d6d96e3e0d39bdbae2c56c](https://github.com/ruebenramirez/.dotfiles/commit/a18fd396a4f53238c7d6d96e3e0d39bdbae2c56c)
