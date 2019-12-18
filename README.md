# TODO List in HTML Forms

This repo demos the usage of HTML forms with a web-based TODO list.

## System Requirements

* [Golang](https://golang.org/)
* [GCC](https://gcc.gnu.org/) for SQLite
* A modern browser to run the app

For Windows users, install GCC provided by [MSYS2](https://www.msys2.org/).

## Usage

### Install the Dependencies

Run *install* (for Unix) or *install.bat* (for Windows) to install the dependencies of the web app.

### Build It

Run *build* (for Unix) or *build.bat* (for Windows) to build the app.

### Run It

Invoke *app* (for Unix) or *app.exe* (for Windows) to run the app.

By default, the app will run on http://127.0.0.1:8080/ . Use `-h` (host) and `-p` (port) to adjust its URL.

The app stores its data into an in-memory SQLite database. Therefore, no persistent data will be generated after the app ends.

### Clean It

Run *clean* (for Unix) or *clean.bat* (for Windows) to remove the app.

## Copyright

2019, Michael. This repo is licensed under [MIT](https://opensource.org/licenses/MIT).
