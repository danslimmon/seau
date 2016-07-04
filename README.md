`seau` is a (proof of concept) metric collector that preserves distribution data through histograms.

Data is fed in with the [statsd](https://github.com/etsy/statsd) protocol. As the `seau-collect` process reads data, it populates a [high dynamic range histogram](https://github.com/HdrHistogram/HdrHistogram), which preserves information about the data's distribution at the cost of precision. In other words, instead of keeping track of the *average* or *sum* of values (or the original values themselves), `seau-collect` keeps track of the *count* of values that fall within a given "bucket". After every flush interval, `seau-collect` sends its current histogram to a `seau-write` process.

`seau-write` receives histogram data from `seau-collect` processes either writing it to disk or relaying it to another `seau-write` process.

# Configuration

## `seau-collect` configuration

`seau-collect` is configured with environment variables:

* `SEAU_FLUSH_INTERVAL`: The number of seconds between flushes. This will be the time-axis precision of your time series.
* `SEAU_LISTEN`: The IP address and UDP port to listen on, separated by a colon (`:`). Default is `127.0.0.1:8124`.
* `SEAU_SIGNIFICANT_DIGITS`: The minimum number of significant digits that should be preserved by the histogram buckets.

## `seau-write` configuration

* `SEAU_WRITE_LISTEN`: The IP address and UDP port to listen on, separated by a colon (`:`). Default is `127.0.0.1:2002`.

# Protocol

## `seau-collect` protocol

Data points are reported to `seau-collect` using the [statsd](https://github.com/etsy/statsd) protocol. It's a UDP protocol, and the default port is 8124. The only metric type supported is "histogram", a new type indicated by `|h`. So, for example, a value might be written as:

    response.latency:19|h

Multiple data points can be sent in a single datagram with a newline between them, like

    gorets:1|h\nglork:320|h\ngaugor:333|h

## `seau-write` protocol

A different protocol is used to transmit histogram data to `seau-write` processes. It's a TCP protocol, and the default port is 2002.
