package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
)

func startup( cassandraCluster string, keyspace string ) *gocql.Session {
	cluster := gocql.NewCluster( cassandraCluster )
	cluster.Keyspace = keyspace
	// cluster.ProtoVersion = 0x3
	session, _ := cluster.CreateSession()
	if session == nil {
		fmt.Fprintf(os.Stderr, "couldn't get a session\n")
		os.Exit(1)
	}
	return session
}

func lookup(session *gocql.Session, sentence string) [][]string {
	normalized := normalize(sentence)
	candidates := getCandidates(session, normalized)
	final := filterCandidates(normalized, candidates)
	return final
}
func getCandidates(session *gocql.Session, words []string) [][]string {
	var first, rest, defn string
	rval := make([][]string, 0)
	iter := session.Query(`select first, rest, defn from defs where first in ?`, words).Consistency(gocql.One).Iter()
	for iter.Scan(&first, &rest, &defn) {
		rval = append(rval, []string{first, rest, defn})
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return rval
}
func getRandom(session *gocql.Session) [][]string {
	var first, rest, defn string
	rval := make([][]string, 0)
	random, _ := gocql.RandomUUID()
	iter := session.Query(`select first, rest, defn from defs where TOKEN(first) > TOKEN(?) limit 1`, random.String()).Consistency(gocql.One).Iter()
	for iter.Scan(&first, &rest, &defn) {
		rval = append(rval, []string{first, rest, defn})
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return rval
}
