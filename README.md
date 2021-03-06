<p align="center">
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/release/crazy-max/travis-wait-enhanced.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/crazy-max/travis-wait-enhanced/releases/latest"><img src="https://img.shields.io/github/downloads/crazy-max/travis-wait-enhanced/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://github.com/crazy-max/travis-wait-enhanced/actions?workflow=build"><img src="https://img.shields.io/github/workflow/status/crazy-max/travis-wait-enhanced/build?label=build&logo=github&style=flat-square" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/crazy-max/travis-wait-enhanced"><img src="https://goreportcard.com/badge/github.com/crazy-max/travis-wait-enhanced?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/travis-wait-enhanced"><img src="https://img.shields.io/codacy/grade/2a33c37cd24e4225adacd48736c0efbb.svg?style=flat-square" alt="Code Quality"></a>
  <br /><a href="https://github.com/sponsors/crazy-max"><img src="https://img.shields.io/badge/sponsor-crazy--max-181717.svg?logo=github&style=flat-square" alt="Become a sponsor"></a>
  <a href="https://www.paypal.me/crazyws"><img src="https://img.shields.io/badge/donate-paypal-00457c.svg?logo=paypal&style=flat-square" alt="Donate Paypal"></a>
</p>

## :warning: Abandoned project

This project is not maintained anymore and is abandoned. Feel free to fork and make your own changes if needed.

Thanks to everyone for their valuable feedback and contributions.

## About

**travis-wait-enhanced** :alarm_clock: is a CLI application written in [Go](https://golang.org/) to replace the
existing [travis_wait](https://docs.travis-ci.com/user/common-build-problems/#build-times-out-because-no-output-was-received)
bash function (with some enhancements) to prevent [Travis CI](https://travis-ci.com/) from thinking a long-running
process has stalled.

## Features

* Custom timeout duration
* Custom interval duration at which to print keep-alive messages
* Custom formatting options
* Print command output to stdout/stderr

## Download

travis-wait-enhanced binaries are available in [releases](https://github.com/crazy-max/travis-wait-enhanced/releases)
page.

Choose the archive matching the destination platform and extract travis-wait-enhanced:

```
$ cd /opt
$ wget -qO- https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.2.0/travis-wait-enhanced_1.2.0_linux_x86_64.tar.gz | tar -zxvf - travis-wait-enhanced
```

After getting the binary, it can be tested with `./travis-wait-enhanced --help` or moved to a permanent location.

```
$ ./travis-wait-enhanced --help
Usage: travis-wait-enhanced <command> ...

Prevent Travis CI from thinking a long-running process has stalled. More info:
https://github.com/crazy-max/travis-wait-enhanced

Arguments:
  <command> ...    Command to execute.

Flags:
  --help               Show context-sensitive help.
  --version
  --timeout=20m        Timeout for this command.
  --interval=1m        Interval at which to print keep-alive messages.
  --print-name         Print the name of this tool to identify keep-alive messages.
  --print-string="Still running..."
                       Keep-alive message printed in each interval.
  --print-timestamp    Print the current timestamp after each keep-alive message.
  --print-newline      Print a newline character after each keep-alive message.
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
    wget -qO- "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.2.0/travis-wait-enhanced_1.2.0_linux_x86_64.tar.gz" | tar -zxvf - travis-wait-enhanced
    mv travis-wait-enhanced /home/travis/bin/
    travis-wait-enhanced --version
  # windows
  - |
    curl -LfsS -o /tmp/travis-wait-enhanced.zip "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v1.2.0/travis-wait-enhanced_1.2.0_windows_x86_64.zip"
    7z x /tmp/travis-wait-enhanced.zip -y -o/usr/bin/ travis-wait-enhanced.exe -r
    travis-wait-enhanced --version
```

## How can I help?

All kinds of contributions are welcome :raised_hands:! The most basic way to show your support is to star :star2:
the project, or to raise issues :speech_balloon: You can also support this project by
[**becoming a sponsor on GitHub**](https://github.com/sponsors/crazy-max) :clap: or by making a
[Paypal donation](https://www.paypal.me/crazyws) to ensure this journey continues indefinitely! :rocket:

Thanks again for your support, it is much appreciated! :pray:

## License

MIT. See `LICENSE` for more details.
