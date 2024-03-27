# gURL

gURL is a Golang command line tool inspired by `curl`. It can be used to make HTTP requests to a server.

## Purpose of the Project

This project was created to provide a simple CLI tool for making HTTP requests. While it's not as comprehensive as `curl`, it serves as a practical introduction to creating CLI tools with Go. This project was also a response to a [coding challenge](https://codingchallenges.fyi/challenges/challenge-curl).

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
