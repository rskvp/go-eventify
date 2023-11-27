# Eventify

A lightweight low-code platform for running workflows on tiny servers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- [Go](https://go.dev/doc/install) `v1.21.4`

### Steps

Clone the repository

```
git clone git@github.com:assalielmehdi/go-eventify.git
cd go-eventify
```

And install dependencies

```
go mod download
```

Then start the server

```
go run .
```

If you need to test some changes, you might want to start the server with live reload

```
air
```

The server should be listening on the port configured in [.env](.env) (default: `8080`).

<!-- ## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
``` -->

## Deployment

These instructions will help you deploy `Eventify` with its [GUI](https://github.com/assalielmehdi/go-eventify-ui) using `Docker` in a live system.

### Prerequisites

- [Docker](https://docs.docker.com/engine/install/)

### Steps

Run docker image

```
docker run -d -p 8080:8080 ghcr.io/assalielmehdi/go-eventify:main
```

### Configuration

`Eventify` comes with a set of defaults. You can override the configuration by setting environment variables

```
docker run -d -p PORT:PORT -e KEY1=VAL1 -e KEY2=VAL2... ghcr.io/assalielmehdi/go-eventify:main
```

The following is an exhaustive list of environment variables

| Key | Configuration |
|-----|---------------|
| key #1 | value #1 |
| key #2 | value #2 |
| key #3 | value #3 |



## Built With

- [Gin](https://gin-gonic.com/) - For routing HTTP traffic
- [GORM](https://gorm.io/index.html) - For ORM
- [SQLite](https://www.sqlite.org/index.html/) - As a lightweight embedded database
- [Testify](https://github.com/stretchr/testify) - For unit and intergation testings

<!-- ## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us. -->

<!-- ## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).  -->

## Authors

- **El Mehdi Assali** - *Initial work* - <assalielmehdi@gmail.com>

See also the list of [contributors](https://github.com/assalielmehdi/go-eventify/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
