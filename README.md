#  Social Poll

This a practice project.

 ## Twitter Votes

Reads tweets and pushes the votes into the messaging queue. Twitter Votes pulls the relevant tweet data, figures out what is being votes for (or rather, which options are mentioned), and pushes the vote into NSQ.

## Counter

Listens out for votes on the messaging queue, and periodically saves the results in the MongoDB database. Counter receives the vote messages from NSQ and keeps an in-memory tally of the results, periodically pushing an update to persist the data.

## Web

Web server that exposes the live results.
