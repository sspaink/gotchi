# Gotchi

Your very own digital Gopher companion.

## Getting Started

### How to load this app to a Badger 2040W

1. [Install Go](https://go.dev/dl/)
2. [Install tinygo](https://tinygo.org/getting-started/install/)

Run the following command:

```sh
tinygo flash -target badger2040-w .
```

### Creating images

1. Draw a picture
2. Save it as a png under the assets folder
3. Run `cd tools/convertimages && go run main.go && cd -` to create a bin file for your new picture, the file will be outputed to the `bin` directory
4. Add it to the code using `go:embed`

## Contributing

Any help is welcome!

* Draw some pictures
* Add new fun things for the Gopher to do

Check the Github issues :)

## Useful resources

* [conejoninja/badger2040](https://github.com/conejoninja/badger2040)
