# The Gator CLI tool
Gator, short for blog aggreGATOR is a CLI tool written in Go that allows you to scrape data from your favorite sites and podcast and then save them to Postgres which will in turn allow you to view this scraped data at a time of your choosing straight from your command line!


## How to get started 

To get started, you will need [Postgres](https://www.postgresql.org/) and [Go](https://go.dev/doc/install) installed on your system, and ive provided links to them
Next you need to install Gator using go [install](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
You will also need [SQLC](https://sqlc.dev/) for getting the db and queries into your machine


## Commands to run in the program

Some commands that are available include

`gator register` - This command accepts a name and registers the name in the db

`gator login` - Logs in a registered user. This command also accepts a name

`gator addfeed` - Accepts a Site name and url and adds that site to the db so you can scrape data later from there

`gator agg` - Takes a time and then fetches data from the saved sites when the passed in time elapses

`gator follow` - Accepts a url and marks that you are following them

`gator following` - Shows all the sites that you are following

