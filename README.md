# Stateful Notification System

## Overview

![Architecture](/assets/overview.png)

## Run the gRPC & API services

```
docker-compose up
make dc-createdb
make dc-up
```
To bring up the `grpcsvc` (gRPC service) and the `apisvc` (API service), run `docker-compose up`. Run `make dc-createdb` to create a `notifier_db` followed by a `make dc-up` to migrate the database tables in `db/migrations`.

```
INSERT INTO users (uuid, name, url) VALUES (1, "name1", "url1");
INSERT INTO users (uuid, name, url) VALUES (2, "name2", "url2");
INSERT INTO users (uuid, name, url) VALUES (3, "name3", "url3");
INSERT INTO users (uuid, name, url) VALUES (4, "name4", "url4");

INSERT INTO followers (uuid, user_uuid, follower_uuid) VALUES (1, 1, 2);
INSERT INTO followers (uuid, user_uuid, follower_uuid) VALUES (2, 1, 3);
INSERT INTO followers (uuid, user_uuid, follower_uuid) VALUES (3, 1, 4);
```
Connect to your mysql container at `127.0.0.1:33060`. The default user is the `root` and the password is `password`. Run the queries to insert mock users & followers data. Those data is needed when we make an API call to the `/notify` handler since that involves listing followers of a single user to enqueue the notification events for.

```
POST http://localhost:9090/notify

{
    "user_uuid": "1",
    "message": "a new event just sent to the notify endpoint!"
}
```
Send requests to the `/notify` endpoint at port `9090` with a data like above. You should observe that the notification events have been enqueued in the `notification_events_queue` table.

## Database Schema

![Database Schema](/assets/database_schema_new.png)

## UML Diagram

![UML Diagram](/assets/uml_diagram.png)
