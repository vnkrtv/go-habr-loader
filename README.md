# go-habr-loader

[![Build Status](https://travis-ci.com/LeadNess/go-habr-loader.svg?branch=master)](https://travis-ci.com/LeadNess/go-habr-loader)
### Description

Loads all posts from habr.com to PostgreSQL.

### Installation

- Clone this repository:
  - ```git clone https://github.com/LeadNess/go-habr-loader.git```
- Set settings (such as PostgreSQL connection information) which are stored in 'config/config.json' file:
  - ```nano go-habr-loader/config/config.json```
- Build docker image:
  - ```docker build -t go-habr-loader go-habr-loader```
- Run docker container:
  - ```docker run --name habr-loader go-habr-loader```