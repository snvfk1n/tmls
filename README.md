# tmls

`tmls` is a CLI overlay for `tmux ls` that facilitates selection and attaching of sessions. Navigate active sessions with <kbd>↓</kbd>/<kbd>↑</kbd> or <kbd>j</kbd>/<kbd>k</kbd>, select one via <kbd>enter</kbd>, or cancel with <kbd>ctrl</kbd>+<kbd>c</kbd> or <kbd>q</kbd>.

### Installation

You may build `tmls` on your own (see below) or simply via `go get github.com/jankaszel/tmls`. Alternatively, you may obtain [pre-built binaries on GitHub](https://github.com/jankaszel/tmls/releases) for macOS or Linux.

### Development

You'll need Go (I recommend using a version newer than 1.10) and `make`. Obtain the source via Git or `go get -d github.com/jankaszel/tmls`, navigate to the cloned directory and build with `make`. Running `make install` will invoke `go install`, which will make the binary available to your system.

## License

`tmls` is licensed under GPL.
