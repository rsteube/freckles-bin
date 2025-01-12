# Init

[Freckles] is backed by a [Git] repository located at `~/.local/share/freckles`.

```go
{{#include ../../../cmd/freckles/cmd/init.go:command}}
```

## Clone

Completion for the `--clone` flag is provided by [ActionRepositorySearch].

```go
{{#include ../../../cmd/freckles/cmd/init.go:flagcompletion}}
```

> [Macros] provide a way to _loosely_ share [Actions] between applications.

[Actions]:https://carapace-sh.github.io/carapace/carapace/action.html
[ActionRepositorySearch]:https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/tools/git#ActionRepositorySearch
[Freckles]:https://github.com/rsteube/freckles
[Git]:https://git-scm.com/
[Macros]:https://carapace-sh.github.io/carapace-spec/carapace-spec/macros.html
