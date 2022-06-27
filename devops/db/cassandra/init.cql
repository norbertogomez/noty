CREATE KEYSPACE "noty"
WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 3};

USE noty;
CREATE TABLE notifications(
   ID           uuid PRIMARY KEY,
   name         text,
   description  text,
   strategy     text,
   notify_at    date,
   user_id      uuid
);