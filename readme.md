# Simple CLI Quote Generator

A simple command-line application built with Go and Cobra that fetches random quotes from an API.

## Features

- Fetch a random quote along with the author's name using the `qt` command.
- Built with the Cobra CLI library for easy command-line interface creation.

## Prerequisites

- Go 1.23.6 or later installed on your system.
- Internet connection to fetch quotes from the API.

## Installation

1. Clone the repository:

```sh
git clone https://github.com/mir-asif01/simple_cli_app.git
cd simple_cli_app
```

2. Install Dependencies:

```sh
    go mod tidy
```

3. Build the Application:

```sh
    go build -o quote
```

4. Run:

```sh
    ./quote qt
```

This will fetch a random quote and display it in the terminal.

Output

```sh
    Quote : "The only limit to our realization of tomorrow is our doubts of today."
    Author : Franklin D. Roosevelt
```

## Thank you
