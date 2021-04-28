# tmls

`tmls` is a CLI overlay for `tmux ls` that facilitates selection and attachment of sessions. Navigate active sessions with <kbd>↓</kbd>/<kbd>↑</kbd> or <kbd>j</kbd>/<kbd>k</kbd>, select one via <kbd>enter</kbd>, or exit with <kbd>ctrl</kbd>+<kbd>c</kbd> or <kbd>q</kbd>. You can create a new session in the current working directory by hitting <kbd>Tab</kbd> and entering a session name, and switch back to session selection by hitting <kbd>Tab</kbd> again.

### Installation

You may build `tmls` on your own (see below) or simply install via `go get github.com/jankaszel/tmls`. Alternatively, you may obtain [pre-built binaries on GitHub](https://github.com/jankaszel/tmls/releases) for macOS or Linux.

### Development

You'll need Go (I recommend using a version newer than 1.13). Obtain the source via Git or `go get -d github.com/jankaszel/tmls`, navigate to the cloned directory and build with `go build`. You may run `go install`,  which will make the binary globally available to your system.

## License

`tmls` is licensed under GPL.

