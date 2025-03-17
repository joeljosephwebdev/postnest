# PostNest 🗞️🪺

Create your own cozy corner of the web with PostNest. PostNest is a command-line tool for effortlessly managing and aggregating RSS feeds. 

- **Add & store**: Collect and store RSS feed posts in a PostgreSQL database.
- **Follow/Unfollow**: Follow or unfollow feeds added by other users.
- **Post Summaries**: View aggregated post summaries directly in the terminal with links to full articles.

This is a guided project from Boot.dev, called gator. 

## Authors 🙋‍♂️

- [Joel Joseph](https://www.github.com/joeljosephwebdev)

## Prerequisites 🚀

For this project you will need Go v1.23 or later,  Postgres,  [goose](https://github.com/pressly/goose) and [sqlc](https://docs.sqlc.dev/en/latest/) installed. You can see the versions I used while building this project below.

* check versions
  ```sh
  >> go version
  go1.23.5 darwin/arm64

  >> postgres --version
  postgres (PostgreSQL) 15.12

  >> sqlc version
  v1.28.0

  >> goose -version
  goose version: v3.24.1


## Getting Started 💫

Add the db_url for your db to the postnestconfig.json file.

* migrate database
  ```sh
  ./dbMigrationUp.sh
  2025/03/12 22:11:22 OK   002_feeds.sql (14.37ms)
  2025/03/12 22:11:22 goose: successfully migrated database to version: 2

* generate SQL commands
  ```sh
  sqlc generate

Now the app is ready to use.

## Usage 🧑‍💻

You can operate the app from the root folder.

* reset database
   ```sh
   go run . reset

* register user
  ```sh
  go run . register <username>

* switch user
  ```sh
  go run . login <username>

* list all users
  ```sh
  go run . users

* add new feed for current user
  ```sh
  go run . addfeed <name> <url>

* list all feeds
  ```sh
  go run . feeds

* follow existing feed
  ```sh
  go run . follow <url>

* show current user's feeds
  ```sh
  go run . following

* unfollow feed 
  ```sh
  go run . unfollow <url>

* start the aggregator loop to pull and store feeds to the db
  ```sh
  go run . agg <duration_string>

* browse saved feeds
  ```sh
  go run . browse <limit>[optional]