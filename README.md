# gURL

gURL is a Golang curl command line tool that can be used to make HTTP requests to a server.

## Purpose of the project

The purpose of this project was to create a simple CLI tool inspired by the `curl` command line tool. It is not as complete as `curl` and indeed do not intend to be, but it can be used to make simple HTTP requests to a server.

It was also made to learn more about Golang and how to create CLI tools with it, more specifically on complete this [challenge](https://codingchallenges.fyi/challenges/challenge-curl)

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
