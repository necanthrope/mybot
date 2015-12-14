# sailbot

`sailbot` is an working Slack bot that asks cassandra for a definition of
a word, then responds with the definition. It's used to help teach sailing
terms.

To get this started, spin up a cassandra node and create a keyspace named
'sailbot' with a table named 'words':

    cqlsh> CREATE KEYSPACE sailbot WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};
    cqlsh> use sailbot;
    cqlsh:sailbot> create table words ( word text primary key, defn text);
    cqlsh:sailbot> copy words (word, defn) from 'words.csv';
    1285 rows imported in 0.365 seconds.

Then, to ask sailbot to define a word, have it join a channel (see the
blog post), then say something like @sailbot: define port

Forked from here:
Check the [blog post](https://www.opsdash.com/blog/slack-bot-in-golang.html)
for a description of mybot internals.
