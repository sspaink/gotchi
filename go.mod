module gotchi

go 1.23.2

require (
	github.com/briandowns/spinner v1.23.1
	tinygo.org/x/drivers v0.29.0
	tinygo.org/x/tinyfont v0.5.0
)

require (
	github.com/fatih/color v1.7.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/term v0.28.0 // indirect
)

//replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240515082542-5f2645f5444d
