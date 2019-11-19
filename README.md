# expense-tracker

A REST API written in golang for budgeting and expense tracking

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

If you don't have them yet, install [golang](https://golang.org/dl/) and [sqlite](https://www.sqlite.org/download.html) first.

### Development

Clone the project (preferrably in your `$GOPATH/src/github.com/<your-username>`):

``` sh
git clone https://github.com/desi-belokonska/expense-tracker
```

Install dependencies:

``` sh
cd expense-tracker
go get
```

Seed initial data (optional):

```sh
./run-seed.sh
```

There is a startup script which simplifies the development process:

``` sh
./run-start.sh
```

If everything went right, you should see this in your terminal window:

``` sh
Listening on port: 5000 (http://localhost:5000)
```

You can also set the PORT env variable to something else.

## Running the tests

To run the tests, simply run

``` sh
go test
```

## Authors

* **Desislava Belokonska** - [desi-belokonska](https://github.com/desi-belokonska)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
