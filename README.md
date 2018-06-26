# A golang logging utility library

[![Build Status](https://travis-ci.org/corvus-ch/logr.svg?branch=master)](https://travis-ci.org/corvus-ch/logr)
[![Maintainability](https://api.codeclimate.com/v1/badges/0c85b21a5a91e898a958/maintainability)](https://codeclimate.com/github/corvus-ch/logr/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/0c85b21a5a91e898a958/test_coverage)](https://codeclimate.com/github/corvus-ch/logr/test_coverage)
[![GoDoc](https://godoc.org/github.com/corvus-ch/logr?status.svg)](https://godoc.org/github.com/corvus-ch/logr)

This library contains a collection of utilities and implementations built around
the `logr.Logger` interface by [Brian Ketelsen][bketelsen].

Implementations are provided for:

- [golangs `log.Logger`][log.logger]
- [logrus by Simon Eskildsen][logrus]
- [zap by Uber][zap]
- [zerolog by Olivier Poitrey][zerolog]

Each of the above implementation comes with a benchmark. The short version:
when interested in having controll about the format and destination of the
output, go with logrus. If performance is the main concern, go with zerolog.

There is also in [implementation using an internal buffer][buffered].

The package [log] provides a global logger which aims to be compatible to the
one provided by `log.Logger`.

Sometimes one might want to use a logger through the `io.Writer` interface. This
is where the package [writer_adapter] comes in handy.

## Contributing and license

This library is licenced under [MIT](LICENSE). For information about how to
contribute to this project, see [CONTRIBUTING.md].

[CONTRIBUTING.md]: https://github.com/corvus-ch/logr/blob/master/CONTRIBUTING.md
[bketelsen]: https://github.com/bketelsen
[buffered]: https://godoc.org/github.com/corvus-ch/logr/buffered
[log.logger]: https://godoc.org/github.com/corvus-ch/logr/log
[log]: https://godoc.org/github.com/corvus-ch/logr/log
[logrus]: https://godoc.org/github.com/corvus-ch/logr/logrus
[writer_adapter]: https://godoc.org/github.com/corvus-ch/logr/writer_adapter
[zap]: https://godoc.org/github.com/corvus-ch/logr/zap
[zerolog]: https://godoc.org/github.com/corvus-ch/logr/zerolog
