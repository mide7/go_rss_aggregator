# A simple rss but powerful aggregator thanks to go!

This is just a fun project for me to get started with golang, which I am considering to build the backend services for most of my projects from hereon. Most of the functionality works with a few glitches here and there which I will fix in the future for sure when I have less exciting things happening in life

Anyways , to run this (cd into the source directory) and do a:

-- go run .

A go routine is spawned which releases 10(change this to whatever you want) other goroutines which go out and fetch the feeds. Ticker is set to run every minute, feel free to change this to something else.

But obviously before you do anything , make sure to add users,feeds and feedfollows. Look at main.go which has all the api endpoints you can hit to perform the operations.

This is what I used in this project:

Go , obviously.
chi router, for routing.
sqlc , to compile sql to type-safe go code
goose, for database migrations. and many other amazing libraries from the standard package.
