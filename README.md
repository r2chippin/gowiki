# gowikiEvo

A Refactored Version of [Writing Web Applications - The Go Programming Language](https://go.dev/doc/articles/wiki/) Using Gin, Gorm, MySQL, and HTMX

## Description

This project is a refactored version of the tutorial [Writing Web Applications - The Go Programming Language](https://go.dev/doc/articles/wiki/), which demonstrates how to build web applications using Go. In this refactored version, we have utilized the following technologies and tools:

- [Gin](https://github.com/gin-gonic/gin): A web framework for Go that provides a fast and flexible way to build web applications.
- [Gorm](https://github.com/go-gorm/gorm): A powerful ORM (Object-Relational Mapping) library for Go, which simplifies database operations.
- MySQL: A popular open-source relational database management system.
- [HTMX](https://htmx.org/): A JavaScript library that allows you to create web pages with dynamic and interactive features using HTML attributes.

## Features

- Web interface for creating and editing wiki pages.
- Persistent storage of wiki pages in a MySQL database.
- Real-time updates using HTMX for seamless and interactive user experience.

## Installation

To run this project locally, follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/r2chippin/gowikiEvo.git
   ```

2. Install the required dependencies:

   ```
   go mod download
   ```

3. Set up the MySQL database and configure the connection in the application.

4. Run the application:

   ```
   go run ./cmd/wiki.go
   ```

5. Access the application in your web browser at `http://localhost:8080/wiki/view/<wiki>`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Acknowledgements

This project is a refactored version of the tutorial [Writing Web Applications - The Go Programming Language](https://go.dev/doc/articles/wiki/), which is provided by the Go programming language community.
