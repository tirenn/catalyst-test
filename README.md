# Catalyst Test App

## How To Run

First of all you need to create folder under **_/src_** by run this command

    mkdir tirenn && cd $_

Clone this repository to your machine

    git@github.com:tirenn/catalyst-test.git    

Move to folder **_/catalyst-test_**

    cd catalyst-test

Then copy **_.env.example_** and rename to **_.env_**

    cp .env.example .env

Edit **_.env_** file to your preference then run command

    source .env

Lastly run command

    go run main.go

## DB Migration

First of all you need to install goose, here is how to install goose for database migration

https://github.com/pressly/goose

To migrate database run this command

    goose -dir db/migrations sqlite3 ./test.db up

To rollback migration run

    goose -dir db/migrations sqlite3 ./test.db down

or you can see the [documentation](https://github.com/pressly/goose)

## API Documentation

For the documentation you can see [here](https://documenter.getpostman.com/view/21322745/2s83S89r5R)

## How To Run Test

Move to folder **_/auth-app_** or **_/fetch-app_** then run command

    go test ./domains/...
