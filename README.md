# Process

A simple Go module designed to make running external processes easier.

## Installation

```bash
go get github.com/jasonmccallister/process
```

## Usage

```go
package main

func main() {
    // start one process
    go func() {
        if err := process.Start(process.Options{
            Name: "echo",
            Args: []string{"Hello"},
            Writer: os.Stdout,
            ErrWriter: os.Stderr,
        }); err != nil {
            log.Fatal(err)
        }
    }()

    // start another
    go func() {
        if err := process.Start(process.Options{
            Name: "echo",
            Args: []string{"Hello"},
            Writer: os.Stdout,
            ErrWriter: os.Stderr,
        }); err != nil {
            log.Fatal(err)
        }
    }()

    // another blocking call such as a web server
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```
