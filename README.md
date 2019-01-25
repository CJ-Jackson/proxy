# Proxy

Precompiled proxy server built for the company I currently work for
(feelunique)

It's for using a dev site on real mobile devices.

It's based on the work of https://github.com/elazarl/goproxy, but 
there is no precompiled builds in that repository and not everyone
has go installed on their machine.

## Installation

Downloads latest release from https://github.com/CJ-Jackson/proxy/releases

Copy the build to `/usr/local/bin` and make sure the
permission is `755` (chmod) and the owner is `root` (chown)

## Usage

```sh
$ proxy
```

and set the browser proxy to `127.0.0.1:8888`

## Note

There is no ip restriction.