# Noty

This project is a simple notification service,
which creates notifications applying different strategies for multiple users.

This project is part of a playground, where I try things that might or might not make sense.


### Prerequisites for running 📋

_Dependencies (minimum tested versions)_

* Docker version 20.10.7
* docker-compose version 1.29.2
* Make
* golang 1.18

### Set up 🔧

To add make commands

## Built with 🛠️

* [Golang](https://golang.org/) - The language
* [Cassandra](https://cassandra.apache.org/_/index.html) - Notifications DB
* [MySQL](https://www.mysql.com/) - Users DB

## Authors ✒️

* **Norberto Gómez** - *Owner* - [norbertogomez](https://github.com/norbertogomez)

## TODOs

- [ ] Add metrics system
- [ ] Add monitoring
- [ ] Add alerting
- [ ] Notification strategies like: Slack, email, whatsapp, etc.
- [ ] Notification pulling
- [ ] Add protocol buffers
- [ ] Add testing
- [ ] Automate start up
- [ ] Add kubernetes set up
- [ ] Set up event based communication with Kafka
- [ ] Play around with caching
- [ ] Add a logging tool like pushing logs to ES or similar
- [ ] Add gRPC