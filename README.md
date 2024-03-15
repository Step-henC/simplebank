### Simple Bank 


## What is This?
    API backend service for simple bank with.



# Local Development Set up

    Please have the following dependencies installed

        - Docker for Postgres DB
        - Make for simplifying cli commands
        - SQLC for generating DB models and queries
        - Migrate [for populating local database](https://github.com/golang-migrate/migrate) 

    I suggest using all software on a linux distro 
    If on Windows and need linux:
        Download Ubuntu (WSL)
        Open cli
        Type: "explorer.exe . " Include the final dot
        You should see a file  system gui
        You can drag and drop this cloned project into the explorer.exe Linux gui file system


# Running Project locally

    The Makefile contains the list of commands to run. For simplicity, run in the cli:
        - make postgres
        - make createdb
        - make migrateup (add tables to db)
        - make migrateup1 (add users table)
        - make sqlc (generate golang models)
        - make server (to run api server)
        - make mock (run mock tests)

    Once running: 
        Test api at using postman at localhost:8080 
