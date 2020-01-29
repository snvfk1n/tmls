# tmls

`tmls` is a CLI overlay for `tmux ls` that facilitates selection and attaching of sessions. Navigate active sessions with <kbd>↓</kbd> and <kbd>↑</kbd>, select one by hitting <kbd>enter</kbd>, or cancel with <kbd>ctrl</kbd>+<kbd>c</kbd>.

### Development

You'll need Go (I recommend using 1.10) and `make`. Get the source via `go get github.com/fallafeljan/tmls`, navigate to `$GOPATH/src/github.com/fallafeljan/tmls` and build with `make`. Running `make install` will trigger `go install`, which makes the binary available to your system.
