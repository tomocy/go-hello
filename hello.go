package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Stdout, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(w io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	addr := flags.String("addr", ":80", "address to listen and serve")

	if err := flags.Parse(args[1:]); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", sayHello)

	fmt.Fprintf(w, "listen and serve on %s\n", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func sayHello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Ver 1: Hello!\n")
}
