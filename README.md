<p align="center">
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/travis-wait-enhanced.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/travis-wait-enhanced/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://github.com/crazy-max/travis-wait-enhanced/actions"><img src="https://github.com/crazy-max/travis-wait-enhanced/workflows/build/badge.svg" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/travis-wait-enhanced"><img src="https://goreportcard.com/badge/github.com/crazy-max/travis-wait-enhanced?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/travis-wait-enhanced"><img src="https://img.shields.io/codacy/grade/2a33c37cd24e4225adacd48736c0efbb.svg?style=flat-square" alt="Code Quality"></a>
  <br /><a href="https://www.patreon.com/crazymax"><img src="https://img.shields.io/badge/donate-patreon-f96854.svg?logo=patreon&style=flat-square" alt="Support me on Patreon"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-00457c.svg?logo=paypal&style=flat-square" alt="Donate Paypal"></a>
</p>

## About

**travis-wait-enhanced** :alarm_clock: is a CLI application written in [Go](https://golang.org/) to replace the existing [travis_wait](https://docs.travis-ci.com/user/common-build-problems/#build-times-out-because-no-output-was-received) bash function (with some enhancements) to prevent [Travis CI](https://travis-ci.com/) from thinking a long-running process has stalled.

## Features

* Custom timeout duration
* Custom interval duration at which to print keep-alive messages
* Custom formatting options
* Print command output to stdout/stderr

## Download

travis-wait-enhanced binaries are available in [releases](https://github.com/crazy-max/travis-wait-enhanced/releases) page.

Choose the archive matching the destination platform and extract travis-wait-enhanced:

```
$ cd /opt
$ wget -qO- https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.0.0/travis-wait-enhanced_1.0.0_linux_x86_64.tar.gz | tar -zxvf - travis-wait-enhanced
```

After getting the binary, it can be tested with `./travis-wait-enhanced --help` or moved to a permanent location.

```
$ ./travis-wait-enhanced --help
usage: travis-wait-enhanced [<flags>] <command>...

Prevent Travis CI from thinking a long-running process has stalled. More info:
https://github.com/crazy-max/travis-wait-enhanced

Flags:
  --help             Show context-sensitive help (also try --help-long and
                     --help-man).
  --timeout=20m      Timeout for this command.
  --interval=1m      Interval at which to print keep-alive messages.
  --print-name       Print the name of this tool to identify keep-alive
                     messages.
  --print-string="Still running..."
                     Keep-alive message printed in each interval.
  --print-timestamp  Print the current timestamp after each keep-alive message.
  --print-newline    Print a newline character after each keep-alive message.
  --version          Show application version.

Args:
  <command>  Command to execute.
```

## Usage

`travis-wait-enhanced [<flags>] <command>`

If your command contains flags (e.g. `mvn -V install -Pmy-profile`) then use `--` to indicate the end of
the travis-wait-enhanced flags to avoid parsing errors. For example

```
travis-wait-enhanced --interval=1m --timeout=20m -- mvn -V clean install -Prun-its
```

To use travis-wait-enhanced in your `.travis.yml` add :

```yml
before_install:
  # linux
  - |
    wget -qO- "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.0.0/travis-wait-enhanced_1.0.0_linux_x86_64.tar.gz" | tar -zxvf - travis-wait-enhanced
    mv travis-wait-enhanced /home/travis/bin/
    travis-wait-enhanced --version
  # windows
  - |
    curl -LfsS -o /tmp/travis-wait-enhanced.zip "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.0.0/travis-wait-enhanced_1.0.0_windows_x86_64.zip"
    7z x /tmp/travis-wait-enhanced.zip -y -o/usr/bin/ travis-wait-enhanced.exe -r
    travis-wait-enhanced --version
```

## How can I help ?

All kinds of contributions are welcome :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Support me on Patreon](.res/patreon.png)](https://www.patreon.com/crazymax) 
[![Paypal Donate](.res/paypal.png)](https://www.paypal.me/crazyws)

## License

MIT. See `LICENSE` for more details.
