# all

**all** is a set of common functions that I need sometimes.

These functions range from sorting slices and maps to various functional
programming patterns. Documentation is available on [pkg.go.dev](docs).

The library is split into mutiple different packages, making it easier to
only import the ones I need, under a nicer namespace.

Because this looks a lot better:

```go
import "github.com/qeaml/all/slices"

func something(slice []int) {
  slices.SortAsc(slice)
}
```

Than this:

```go
import "github.com/qeaml/all"

func something(slice []int) {
  all.SortSliceAsc(slice)
}
```

You can of course rename the namespace in case it conflicts with a package
already present in your code:

```go
import qslices "github.com/qeaml/all/slices"

func something(slice []int) {
  qslices.SortAsc(slice)
}
```

## License

This library is licensed under the [BSD 3-Clause][bsd3] license. You can view
it in the [`LICENSE`](LICENSE) file.

[docs]: https://pkg.go.dev/github.com/qeaml/all
[bsd3]: https://opensource.org/licenses/BSD-3-Clause
