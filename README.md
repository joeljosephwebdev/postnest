# PostNest 🗞️🪺

Create your own cozy corner of the web with PostNest. PostNest is a command-line tool for effortlessly managing and aggregating RSS feeds. 

- **Add & store**: Collect and store RSS feed posts in a PostgreSQL database.
- **Follow/Unfollow**: Follow or unfollow feeds added by other users.
- **Post Summaries**: View aggregated post summaries directly in the terminal with links to full articles.

## Getting Started 💫

* migrate database
  ```sh
  ./dbMigrationUp.sh
  2025/03/12 22:11:22 OK   002_feeds.sql (14.37ms)
  2025/03/12 22:11:22 goose: successfully migrated database to version: 2

* generate SQL commands
  ```sh
  sqlc generate

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


