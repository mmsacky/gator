# Gator

Gator is a multi-user CLI RSS feed aggregator inspired by applications like Feedly and Google Reader.

It allows multiple users to add and follow RSS feeds, browse posts, and automatically fetch new content at a configurable interval.

Features

* Register and manage multiple users
* Switch between users
* View all registered users
* Add RSS feeds
* Follow and unfollow feeds added by other users
* View feeds you’ve added or followed
* See which user originally added each feed
* Automatically fetch new posts at a configurable interval
* Browse collected posts

Requirements

* Go 1.27.1
* PostgreSQL
* Goose
* SQLC

Installation

1. Install Go

If you don’t already have Go installed, follow the official installation guide:

https://go.dev/doc/install

2. Install PostgreSQL

macOS

brew install postgresql@15

Linux / WSL (Debian)

sudo apt update
sudo apt install postgresql postgresql-contrib

3. Install Goose

go install github.com/pressly/goose/v3/cmd/goose@latest

4. Install SQLC

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

5. Install Gator

go install github.com/mmsacky/gator@latest

Configuration

Create a configuration file at:

~/.gatorconfig.json

Add your PostgreSQL connection string:

{
  "db_url": "postgres://username:@localhost:5432/gator?sslmode=disable"
}

Replace username with your PostgreSQL username.

Database Setup

Start PostgreSQL.

macOS

brew services start postgresql@15

Linux

sudo service postgresql start

Create the gator database if it doesn’t already exist:

createdb gator

Then run the database migrations:

goose postgres postgres://username:@localhost:5432/gator up

Replace the connection string with the same one used in your ~/.gatorconfig.json.

Usage

Register a User

Register your first user:

gator register "your_username"

You can also run the application directly from the project directory:

go run . register "your_username"

Start the Aggregator

In a separate terminal, start the feed aggregator:

gator agg 1m

The interval can be specified using standard Go duration values:

1s
1m
1h

For example:

gator agg 30s

The aggregator will periodically fetch new posts from the RSS feeds in the database.

Commands

Command	Description
register	Register a new user
login	Set the current user
users	List all registered users
reset	Reset the database to a clean state
addfeed	Add a new RSS feed
feeds	List all RSS feeds
follow	Follow an RSS feed
following	List feeds followed by the current user
unfollow	Unfollow an RSS feed
agg	Start the RSS feed aggregator
browse	Browse posts collected from RSS feeds

Example Workflow

Register a user:

gator register michael

Add a feed:

gator addfeed "Hacker News" "https://hnrss.org/frontpage"

Start the aggregator:

gator agg 1m

Follow a feed:

gator follow https://hnrss.org/frontpage

View the feeds you’re following:

gator following

Browse collected posts:

gator browse

License

None.