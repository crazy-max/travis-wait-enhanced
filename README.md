<p align="center">
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/travis-wait-enhanced.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/travis-wait-enhanced/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://travis-ci.com/crazy-max/travis-wait-enhanced"><img src="https://img.shields.io/travis/com/crazy-max/travis-wait-enhanced/master.svg?style=flat-square" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/travis-wait-enhanced"><img src="https://goreportcard.com/badge/github.com/crazy-max/travis-wait-enhanced?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/travis-wait-enhanced"><img src="https://img.shields.io/codacy/grade/2a33c37cd24e4225adacd48736c0efbb.svg?style=flat-square" alt="Code Quality"></a>
  <br /><a href="https://www.patreon.com/crazymax"><img src="https://img.shields.io/badge/donate-patreon-fb664e.svg?style=flat-square" alt="Support me on Patreon"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

**travis-wait-enhanced** :alarm_clock: is a CLI application written in [Go](https://golang.org/) to replace the existing [travis_wait](https://docs.travis-ci.com/user/common-build-problems/#build-times-out-because-no-output-was-received) bash function (with some enhancements) to prevent [Travis CI](https://travis-ci.com/) from thinking a long-running process has stalled.

## Features

* Custom timeout duration
* Custom interval duration at which to print keep-alive messages
* Print command output to stdout/stderr

## Download

travis-wait-enhanced binaries are available in [releases](https://github.com/crazy-max/travis-wait-enhanced/releases) page.

Choose the archive matching the destination platform and extract travis-wait-enhanced:

```
$ cd /opt
$ wget -qO- https://github.com/crazy-max/travis-wait-enhanced/releases/download/v0.1.1/travis-wait-enhanced_0.1.1_linux_x86_64.tar.gz | tar -zxvf - travis-wait-enhanced
```

After getting the binary, it can be tested with `./travis-wait-enhanced --help` or moved to a permanent location.

```
$ ./travis-wait-enhanced --help
usage: travis-wait-enhanced [<flags>] <command>...

Prevent Travis CI from thinking a long-running process has stalled. More info:
https://github.com/crazy-max/travis-wait-enhanced

Flags:
  --help         Show context-sensitive help (also try --help-long and
                 --help-man).
  --timeout=20m  Timeout for this command.
  --interval=1m  Interval at which to print keep-alive messages.
  --version      Show application version.

Args:
  <command>  Command to execute.
```

## Usage

`travis-wait-enhanced [<flags>] <command>`

* `--help` : Show help text and exit. _Optional_.
* `--version` : Show version and exit. _Optional_.
* `--timeout <duration>` : Timeout for this command. _Optional_. (default: `20m`).
* `--interval <duration>` : Interval at which to print keep-alive messages. _Optional_. (default: `1m`).

## How can I help ?

All kinds of contributions are welcome :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Support me on Patreon](.res/patreon.png)](https://www.patreon.com/crazymax) 
[![Paypal Donate](.res/paypal-donate.png)](https://www.paypal.me/crazyws)

## License

MIT. See `LICENSE` for more details.
