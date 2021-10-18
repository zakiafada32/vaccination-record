# Vaccination Record

## Overview

- ### ERD Diagram

![ERD diagram](https://raw.githubusercontent.com/zakiafada32/vaccination-record/main/ERD.PNG)

- [ERD diagram](https://github.com/zakiafada32/vaccination-record/blob/main/ERD.PNG)

- ### API Documentation

  - [API Documentation](https://github.com/zakiafada32/vaccination-record/blob/main/vaccination-record.postman_collection.json)

## Run locally

- ### Prerequisites

  To get started, ensure that you have docker installed on your local machine.

  ```
  $ docker version
  ```

- ### build and run the app

  To build and run the app with docker-compose command

  ```
  $ docker-compose up --build
  ```

- ### Create database table
  ```
  $ docker exec -i vaccination-mysql sh -c 'exec mysql -uroot -p"$MYSQL_ROOT_PASSWORD" vaccine' < migrate.sql
  ```
