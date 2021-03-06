package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

var expectedSize int

type rewriter func(io.Writer, io.Reader) error

// RunRewriters runs all the given rewriters on the origin reader
func RunRewriters(rews []rewriter, orig io.Reader) (io.Reader, error) {
	var (
		src bytes.Buffer
		dst bytes.Buffer
		err error
	)
	io.Copy(&src, orig)
	for _, r := range rews {
		// log.Printf("Executing rewriter #%d", i)
		err = r(&dst, &src)
		if err != nil {
			return nil, err
		}
		src.Reset()
		dst, src = src, dst
	}
	// log.Print("Success, returning last buffer!")
	return &src, nil
}

// Rews is a slice of rewriters
var Rews = []rewriter{
	// Fix CRLF madness
	func(w io.Writer, r io.Reader) error {
		s := bufio.NewScanner(r)
		for s.Scan() {
			fmt.Fprintln(w, s.Text())
		}
		return s.Err()
	},
	// Remove parentheses
	func(w io.Writer, r io.Reader) error {
		s := bufio.NewScanner(r)
		for s.Scan() {
			spl := strings.Split(s.Text(), " ")
			for _, r := range spl {
				if strings.Contains(r, "(") {
					r = strings.Split(r, "(")[0]
				}
				w.Write([]byte(r))
				w.Write([]byte{' '})
			}
			w.Write([]byte{'\n'})
		}
		return s.Err()
	},
	// Remove "H.C."
	func(w io.Writer, r io.Reader) error {
		s := bufio.NewScanner(r)
		for s.Scan() {
			r := strings.Replace(s.Text(), " H.C. ", " ", -1)
			w.Write([]byte(r))
			w.Write([]byte{'\n'})
		}
		return s.Err()
	},
	// Remove double spaces
	func(w io.Writer, r io.Reader) error {
		s := bufio.NewScanner(r)
		for s.Scan() {
			r := strings.Replace(s.Text(), "  ", " ", -1)
			w.Write([]byte(r))
			w.Write([]byte{'\n'})
		}
		return s.Err()
	},
	// Fix newlines inbetween words
	func(w io.Writer, r io.Reader) error {
		s := bufio.NewScanner(r)
		prev := ""
	scanning:
		for s.Scan() {
			text := strings.TrimSpace(s.Text())
			if (strings.Count(text, " ") + 1) >= expectedSize {
				w.Write([]byte{'\n'})
				w.Write([]byte(text))
				prev = text
				continue
			}
			log.Printf("Found short line: %q, exp size %d", text, expectedSize)
			log.Printf("Count of words is: %d", (strings.Count(text, " ") + 1))
			combined := prev + " " + text
			// TODO: Replace cities with last column
			for _, c := range doubleNameCities {
				if !strings.HasSuffix(combined, c) {
					//log.Printf("%q doesn't end with %q", combined, c)
					continue
				}
				log.Printf("Found city: %s", c)
				w.Write([]byte{' '})
				w.Write([]byte(text))
				prev = "" // I don't want to perform double joins
				continue scanning
			}

			// Sometimes name and surname are glue together and joining below won't work
			for _, c := range cities {
				if strings.Contains(combined, c) {
					log.Printf("Found city %s, in line %s. Not going to join below.", c, combined)
					w.Write([]byte{'\n'})
					w.Write([]byte(text))
					continue scanning
				}
			}

			w.Write([]byte{'\n'})
			w.Write([]byte(text))

			s.Scan()
			// Joining below only if it is not the last line of the file
			if s.Text() != "" {
				w.Write([]byte{' '})

				prev = text + " " + s.Text()
				w.Write([]byte(s.Text()))

				log.Printf("Joining below: '%q, result: '%q'", text, prev)
			}
		}
		return s.Err()
	},
}
