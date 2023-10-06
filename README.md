serve
=====

Serve is a HTTP file server command. Every Go developer has built one, and this
one is mine.

It serves up a directory of content for you to play with in a browser or other
HTTP client. Requests to `/my/path` will be translated to
`/passed/in/dir/my/path`.

When I'm working on the webserver we built at work, I like to use it as the
backend service for testing requests with `-v` turned on.

By default, serve will boot on `localhost:10000` with the current directory as
its serving path.

Usage
-----

    $  serve -h
    usage: serve [DIR]
    -H value
            colon-separated header key and value to set on the response, may be specified multiple times
    -addr string
            address to serve from (default "localhost:10000")
    -q	log nothing
    -v	verbose logging, dumping requests

Installing
----------

The easiest way of installing serve is to [install Go][installgo] (being sure
to set up a working `$GOPATH`, detailed in those instructions), and running

    go get github.com/jmhodges/serve

[installgo]: http://golang.org/doc/install#install
