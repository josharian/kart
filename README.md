# (Go) Kart

Kart is a simple speed-typing race. It is meant to be used
in demos / workshops. Kart was designed to show a few
interesting aspects of Go, particularly concurrency and interfaces.

Suggestions for further experimentation and alteration (ranging from
small to large) can be found in [more_fun.md](more_fun.md).

To run (and hack on) Kart:

* [Install Go](http://golang.org/doc/install)
* Get the code: `go get github.com/josharian/kart`
* `cd $GOPATH/src/github.com/josharian/kart`
* Run the tests: `go test ./...` or, with coverage, `go test -cover ./...`
* Run go vet: `go vet ./...`
* Format the code: `gofmt -w -s .`
* Build: `go build *.go`
* Run: `./kart`

[On your mark...](http://golang.org/doc/gopher/pencil/gophermega.jpg)

[Get set...](http://golang.org/doc/gopher/pencil/gopherswrench.jpg)

[Go!](http://golang.org/doc/gopher/pencil/gopherrunning.jpg)
