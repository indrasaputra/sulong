# Sulong

[![Go Report Card](https://goreportcard.com/badge/github.com/indrasaputra/sulong)](https://goreportcard.com/report/github.com/indrasaputra/sulong)
[![Workflow](https://github.com/indrasaputra/sulong/workflows/Test/badge.svg)](https://github.com/indrasaputra/sulong/actions)
[![codecov](https://codecov.io/gh/indrasaputra/sulong/branch/main/graph/badge.svg?token=tVuz2Rkgna)](https://codecov.io/gh/indrasaputra/sulong)
[![Maintainability](https://api.codeclimate.com/v1/badges/a4802a34897d6c7f3a71/maintainability)](https://codeclimate.com/github/indrasaputra/sulong/maintainability)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=indrasaputra_sulong&metric=alert_status)](https://sonarcloud.io/dashboard?id=indrasaputra_sulong)
[![Go Reference](https://pkg.go.dev/badge/github.com/indrasaputra/sulong.svg)](https://pkg.go.dev/github.com/indrasaputra/sulong)

## Description

Sulong is a simple application that crawls [TaniFund](https://tanifund.com/)'s projects, checks if there is new project, then notifies user about the new project (if any) in [Telegram](https://telegram.org/) channel.

## Owner

[Indra Saputra](https://github.com/indrasaputra)

## Prerequisites

To use this application, make sure you have a [Telegram Bot](https://telegram.org/faq#bots). 
To avoid DDoS, the application will run once in an hour by default. It can be changed via env configuration.

## How to Run

- Install Go

    This project uses version 1.16. Follow [Golang installation guideline](https://golang.org/doc/install).

- Create your Telegram bot

    Read [https://telegram.org/blog/bot-revolution](https://telegram.org/blog/bot-revolution).

    Follow [https://core.telegram.org/bots](https://telegram.org/blog/bot-revolution) for developer guide.

- Clone the project (use one of the two methods below)

    Use SSH
    ```
    $ git@github.com:indrasaputra/sulong.git
    ```
    
    Use HTTP
    ```
    $ https://github.com/indrasaputra/sulong.git
    ```

- Go to project folder

    Usually, it would be
    ```
    $ cd go/src/github.com/indrasaputra/sulong
    ```

- Fill in the environment variables

    Copy the sample env file.
    ```
    $ cp env.sample .env
    ```
    Then, fill the values according to your setting in `.env` file.

- Download the dependencies

    ```
    $ make tidy
    ```
    or run this command if you don't have `make` installed in your local.
    ```
    $ GO111MODULE=on go mod download 
    ```

- Run the application

    ```
    $ go run cmd/main.go
    ```

## Deployment

Currently, this project is deployed in one of my server. These are the steps I do for deployment:

- Build Go binary

    ```
    $ make compile
    ```

- Copy compiled binary to server

    ```
    $ scp <source> <target>
    ```

    Change `<source>` and `<target>` according to your environment. You may use identity to access the server. For example:

    ```
    $ scp -i ~/.ssh/myserver sulong indra@172.0.0.1:/home/indra
    ```

- SSH to server

    ```
    $ ssh -i ~/.ssh/myserver indra@172.0.0.1
    ```

- Run binary

    ```
    $ TELEGRAM_RECIPIENT_ID=1 TELEGRAM_URL=url TELEGRAM_TOKEN=token TANIFUND_URL=url SLEEP=60 ./sulong
    ```