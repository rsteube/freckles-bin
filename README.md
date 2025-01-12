# freckles

[![PkgGoDev](https://pkg.go.dev/badge/github.com/rsteube/freckles)](https://pkg.go.dev/github.com/rsteube/freckles)
[![GoReportCard](https://goreportcard.com/badge/github.com/rsteube/freckles)](https://goreportcard.com/report/github.com/rsteube/freckles)
[![Coverage Status](https://coveralls.io/repos/github/rsteube/freckles/badge.svg?branch=master)](https://coveralls.io/github/rsteube/freckles?branch=master)

A simple dotfile manager based on [carapace](https://github.com/rsteube/carapace) using the symlink approach.

[![asciicast](https://asciinema.org/a/499658.svg)](https://asciinema.org/a/499658)

## Completion

```sh
# bash
source <(freckles _carapace)

# elvish
eval (freckles _carapace|slurp)

# fish
freckles _carapace | source

# oil
source <(freckles _carapace)

# nushell
freckles _carapace | save freckles.nu ; nu -c 'source freckles.nu'

# powershell
freckles _carapace | Out-String | Invoke-Expression

# tcsh
eval `freckles _carapace`

# xonsh
exec($(freckles _carapace))

# zsh
source <(freckles _carapace)
```
