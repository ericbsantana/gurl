# GURL

GURL is a Golang curl command line tool that can be used to make HTTP requests to a server.

## Installation

To install this project, you need to have Go installed on your machine. Then, you can clone this repository and build the project:

```sh
git clone https://github.com/ericbsantana/gurl
cd gurl
go build
```

## Usage

To use this project, you can run the built binary with the desired options. The following are some examples of how to use this CLI:

```bash
gurl http://eu.httpbin.org/get
gurl http://eu.httpbin.org/bearer -H 'Authorization: Bearer guineapig'
gurl http://eu.httpbin.org/post -X POST -d '{"name": "Robert J. Oppenheimer"}' -H "Content-Type: application/json"
gurl http://eu.httpbin.org/put -X PUT -d '{"name": "Ludwig Wittgenstein"}' -H "Content-Type: application/json"
```

## Contributing

Contributions are welcome. Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
