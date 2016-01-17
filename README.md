# sailbot

`sailbot` is an working Slack bot that asks cassandra for a definition of
a word or phrase, then responds with the definition. It's used to help teach sailing
terms, found in the words.csv file.

Mostly I did this project so I could use the go Cassandra interface. Plus it was
a challenge to do embedded phrase searching within a string efficiently. Thanks to
Cobi Carter for providing some great insight.

To get this started, spin up a cassandra node and create a keyspace named
'sailbot' with a table named 'words':

    cqlsh> CREATE KEYSPACE sailbot WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};
    cqlsh> use sailbot;
    cqlsh:sailbot> create table defs ( first text, rest text, defn text, primary key (first, rest));
    cqlsh:sailbot> copy defs from 'words.csv' with null='*' and escape='"';
    1285 rows imported in 0.365 seconds.

Then, to ask sailbot to define a word, have it join a channel (see the
blog post), then say something like @sailbot: define port

Forked from here:
Check the [blog post](https://www.opsdash.com/blog/slack-bot-in-golang.html)
for a description of mybot internals.
