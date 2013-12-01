serve
=====

Serve is yet another file server command. Every Go developer has built one, and
this one is mine.

It serves up a directory of content for you to play with in a browser or other
HTTP client. Requests to `/my/path` will be translated to
`/passed/in/dir/my/path`.

When I'm working on the webserver we built at work, I like to use it as the
backend service for testing requests with `-vv` turned on.

By default, serve will boot on `localhost:10000` with the current directory as
its serving path.

Usage
-----

    $  serve -h
    usage: serve [DIR]
      -addr="localhost:10000": address to serve from
      -v=false: verbose logging
      -vv=false: very verbose logging, dumping requests

Installing
----------

The easiest way of installing serve is to [install Go][installgo] (being sure
to set up a working `$GOPATH`, detailed in those instructions), and running

    go get github.com/jmhodges/serve

[installgo]: http://golang.org/doc/install#install
