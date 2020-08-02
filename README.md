# go-habr-loader

[![Build Status](https://travis-ci.com/LeadNess/go-habr-loader.svg?branch=master)](https://travis-ci.com/LeadNess/go-habr-loader)
### Description

Loads all posts from habr.com to PostgreSQL.

### Installation

- Clone this repository:
  - ```git clone https://github.com/LeadNess/go-habr-loader.git```
- Set settings (such as PostgreSQL connection information) which are stored in 'config/config.json' file:
  - ```nano config/config.json```
- Install requirements:
  - ```go get -t github.com/anaskhan96/soup github.com/jmoiron/sqlx github.com/lib/pq```
- Run app:
  - ```go run cmd/main.go 2> logs.txt &```