module gotchi

go 1.23.2

require (
	tinygo.org/x/drivers v0.28.0
	tinygo.org/x/tinyfont v0.3.0
)

require github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect

replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240515082542-5f2645f5444d
