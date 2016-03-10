# Bancho [![Build Status](https://travis-ci.org/bnch/bancho.svg?branch=master)](https://travis-ci.org/bnch/bancho) [![Go Report Card](https://goreportcard.com/badge/github.com/bnch/bancho)](https://goreportcard.com/report/github.com/bnch/bancho)

This repository is a clone of the bancho server. Bancho is the software that
powers [osu!](https://osu.ppy.sh). Currently players of osu! are getting a bit
sick of the whole ecosystem, so there are many attempts to move out of it, and
that means that plenty of private servers are popping up, and being developed.

This server is written with the help of:

1. Coffee
2. justm3/HoLLy-HaCKeR's [custom-bancho](https://github.com/HoLLy-HaCKeR/custom-bancho).
3. Last but not least, the private server I and a friend of mine had been
   writing since about August 2015. The server is currently closed source due to
   some privacy issues and bad code in general. That friend of mine in about
   January 2016 started developing the version 1.5 of the said private server,
   which included a bancho server. It's even worse in code than justm3's
   custom-bancho, so don't expect to understand much if you haven't dug deep
   into the structure of the osu! packets. Anyway, if you want to take it for
   reference for making your bancho server, [there you go](http://hastebin.com/opadinohej.php).

## Setting up

For setting up bancho, the only things that we require you to have are a brain
and MySQL. Nothing else must be installed. In the future we might as well
implement SQLite, so that you won't even need mysql to get up and running.

1. Grab the latest [build artifact](http://zxq.co:60291/view/bnch/bancho), and
   move the file to a folder where it's safe to break some stuff if required.
2. Open up a command line and run the executable (with an eventual `chmod +x`).
   If you're on windows it might as well work double clicking, but we're not
   entirely sure our braindead system works.
3. Edit bancho.ini with a text editor
4. Run the executable again
5. Add what's below in that big gray box with those 127.0.0.1 things to your
   hosts file.
6. Connect to osu.ppy.sh in your browser, and sign up.
7. Open up stable fallback and then... profit?

```
127.0.0.1 osu.ppy.sh
127.0.0.1 a.ppy.sh
127.0.0.1 c.ppy.sh
127.0.0.1 c1.ppy.sh
```

