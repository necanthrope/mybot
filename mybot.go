/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	session := startup()
	defer session.Close()
	answers := lookup(session, "what's an anchor chain?")
	if len(answers) != 1 {
		fmt.Printf("%q", answers)
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, id := slackConnect(os.Args[1])
	fmt.Println("mybot ready, ^C exits")

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		// see if we're mentioned
		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			// if so try to parse if
			ans := lookup(session, m.Text)
			if len(ans)>0 {
				// looks good, get the quote and reply with the result
				go func(m Message) {
					for _, def := range ans {
						if len(def[1]) > 0 {
							m.Text = "*" + def[0] + " " + def[1] + "*: " + def[2]
						} else {
							m.Text = "*" + def[0] + "*: " + def[2]
						}
						postMessage(ws, m)
						}
				}(m)
				// NOTE: the Message object is copied, this is intentional
			} else {
				// huh?
				m.Text = fmt.Sprintf("sorry, that does not compute\n")
				postMessage(ws, m)
			}
		}
	}
}

func getDefinition(session *gocql.Session, words []string) string {
	var defn string
	thingtodefine := strings.ToLower(strings.Join(words, " "))
	iter := session.Query("select defn from words where word = ?", thingtodefine).Consistency(gocql.One).Iter()
	for iter.Scan(&defn) {
		return fmt.Sprintf("'%s': %s", thingtodefine, defn)
	}
	return fmt.Sprintf("Sorry I don't know about '%s'", thingtodefine)
}
