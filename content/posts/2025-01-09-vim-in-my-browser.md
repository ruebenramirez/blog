---
title: "Vim in My Browser"
date: 2025-01-09T14:35:04-06:00
categories:
- tech
tags:
- tech-tip
---

_**warning:** you're walking into an uber nerd post about vim keybindings_

A long time ago I spent a lot of time learning to use vim, and it's been time well spent!  I definitely feel more productive while writing code and editing text inside of vim or in tools that support vim keybindings.  The problem is that not all of the applications I use on a daily basis are configured to use vim keybindings for editing text.  Today I took a step towards changing that!

### vimium

I've used the [vimium](https://github.com/philc/vimium) browser extension to navigate multiple tabs worth of websites in my browser without leaving the keyboard.  It's been awesome.  I rarely find the need to use a mouse for navigating websites, and when I do it's for very odd exception type things.  So we can scroll up and down with `j` and `k` keys. When we want to click somewhere on the page, we press `f` and the screen is painted with shortcodes.  Then enter the shortcode for the element you want to click.

<img src="/images/screenshot-2025-01-09-18:52:33.png" alt="" width="800" />

### firenvim

For all that vimium does so well, its main shortcoming is that I still have to edit text in HTML textareas without the help of my vim muscle memory.  It's kind of a jarring experience.  I finally took the time to look for solutions to this problem and found the [firenvim](https://github.com/glacambre/firenvim) browser extension.  It supports chrome (and chromium browser variants) as well as firefox.  The general idea here is that focusing on an HTML text element loads up an instance of neovim directly over the textarea.  This is a "real" instance of neovim, with all of your editor customization goodness.  You use that full blown neovim instance to edit the text you want to submit in that HTML text element and then `:wq` to submit the buffer to the page.  This leaves you with a bunch of text entered into the normal HTML text box and then normal browser usage resumes.

<img src="/images/screenshot-2025-01-09-18:48:08.png" alt="" width="800" />


### the setup

The initial setup is pretty straight forward: there's a vim plugin to install, then separately a browser plugin/extension to install, and finally a browser restart is required.

I had a little bit of trouble initially, but found a tip for <C-e> once you focus in the text area to force the plugin to load.


### a wishlist

This is an amazing improvement over entering text without key bindings inside HTML textareas, but there are a few things that would make this even better.

Whenever a textarea html element is focused on, the firenvim extension loads up neovim.  One thing I think would improve on this would be to only load the extension via the <C-e> keyboard shortcut.

I have a habit of using <C-w> to delete the previous word while I'm in insert mode.  <C-w> while editing text in neovim in one of these textareas, closes the browser tab.  It would be nice to override common keybindings inside the neovim session.


