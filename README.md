# go-mamayo

[![MIT LICENSE](https://img.shields.io/github/license/jugesuke/gomamayo)](./LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jugesuke/gomamayo)
[![GoDev](https://pkg.go.dev/badge/github.com/jugesuke/gomamayo)](https://pkg.go.dev/github.com/jugesuke/gomamayo)

Gomamayo provides checking a term if it is _[Gomamayo](https://3qua9la-notebook.hatenablog.com/entry/2021/04/10/220317)_.

<!-- [[日本語版Readmeはこちら]](./README.ja.md) -->

## Getting Started

### Usage

First, import go-mamayo to your project.
```bash
go get github.com/jugesuke/gomamayo
```
```golang
import "github.com/jugesuke/gomamayo"
```

Next, create `*gomamayo.GoMamayo`.

```golang
g, err := gomamayo.Init()
if err != nil {
  panic(err)
}
```

Next, use `Analyze` function.
```golang
result := g.Analyze("ゴママヨサラダ")
```

To get the result if the term is _Gomamayo_, use `result.IsGomamayo`.

```golang
result := g.Analyze("ゴママヨサラダ")
if result.IsGomamayo {
  fmt.Println("This is Gomamayo")
} else {
  fmt.Println("This is not Gomamayo")
}
```
<!-- More information, please read [Document](https://pkg.go.dev/github.com/jugesuke/gomamayo/). -->

## References
- [【論説】ゴママヨおよびそれに付随するさまざまな現象の研究 - 0と1の間でティータイム](https://3qua9la-notebook.hatenablog.com/entry/2021/04/10/220317)
- [ゴママヨ - 用語集 | しなちくシステム](https://thinaticsystem.com/glossary/gomamayo)
