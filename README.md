![📰 Good News app](https://miro.medium.com/max/3840/1*bf7idsl5jS_aic9toUVUWA.png)

# Intentions
**📰Good News app** is a fully open source project created exceptionally for educational purposes. It contains [backend](https://github.com/kanzitelli/good-news-backend) (you can see the response of it if you open this link `https://api.good-news.ggc.team/v1/news/`) written in [Golang](https://golang.org/), iOS & Android mobile apps (which I will publish to App Store and Google Play) made with [Flutter](https://flutter.dev/) and frontend made with [Hummingbird (Flutter for web)](https://github.com/flutter/flutter_web). Links to repositories for code of mobile apps and frontend will be added later when as soon as I finish working on them. My code is not going to be perfect due to the fact that I am not a super professional in any of those topics but I will do my best, promise ✌️ I am a supporter of an idea of sharing knowledge with the world because it is always good to share what you have learnt with those who might struggle with the same issues you faced and by doing it I will learn something new as well. WIN-WIN strategy. 

# Contents
I am in process of writing chapters divided to articles on Medium. Below you will find links to them. They will be updated as I finish working on them.

[Introductory article](https://medium.com/@kanzitelli/good-news-app-golang-flutter-hummingbird-1949f22b299f) - here you will find all the introductory information and links to other chapters as well.

All chapters of a "book":
1. [📰 Good News app. Backend in Golang behind Traefik reverse proxy with https available.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-behind-traefik-reverse-proxy-with-https-available-2165b11467e4) 
2. *[in progress]* 📰 Good News app. Flutter for rapid mobile applications development.
3. *[in progress]* 📰 Good News app. Hummingbird as a promising replacement for frontend frameworks.

And here are articles of current chapter (repository):
1. [Prerequisites & Idea, project and database structure and API endpoints.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-behind-traefik-reverse-proxy-with-https-available-2165b11467e4)
2. [Project creation, go modules & GIN (beautiful framework) integration.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-project-creation-go-modules-gin-integration-237960136b4c)
3. [Colly usage.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-colly-usage-325ded0b5d34)
4. [MongoDB setup using official Golang driver.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-mongodb-setup-using-official-golang-driver-e1994473c8c0)
5. [Running all together locally with Docker and Docker Compose & Traefik v2.0 configuration.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-docker-compose-traefik-v2-0-54545f5590ed)
6. [Publishing to Digital Ocean, Let's Encrypt and DNS Challenge configuration.](https://medium.com/@kanzitelli/good-news-app-backend-in-golang-digital-ocean-lets-encrypt-and-dns-challenge-a8e44d1be7ec)

# Prerequisites
In order to launch this project locally, you have to have [Golang](https://golang.org/) and [Docker](https://docs.docker.com/v17.12/install/) installed on your machine.

**NOTICE:** code in this repository is not a final version of what I have been telling in my articles. For example, if you would like to know how to publish this project on Digital Ocean droplet behind Traefik reverse proxy with https available on your own domain, you will need to follow all articles above. 

# Installation and running steps
First of all, you have to `cd` to where your `$GOPATH` is pointing to on your machine, then `cd` to `$GOPATH/src/github.com/<your_github_or_any_username>`. After that, you will clone this repository, `cd` to it and run it with Docker Compose. Please, make sure that your Docker is launched successfully while following steps below. 
So the steps are:
- `> cd ~/go/src/github.com/kanzitelli/` - might be different from yours depending on your `$GOPATH` and *username*. In order to check where `$GOPATH` is pointing to, please type `go env` in terminal and you will find it in the beginning of the printed list.
- `> git clone https://github.com/kanzitelli/good-news-backend.git`
- `> cd good-news-backend/`
- `> docker-compose build && docker-compose up -d`
- ... wait for some time till whole process is finished
- open a browser window and go to `localhost:6969/v1/news/sources`. You should see nothing, right. Because we need to configure MongoDB as well 🙂
- `> docker exec -it mongo mongo -u "GGCTeamBatr" -p "MySuperSecretPassword" --authenticationDatabase admin` - to open shell of mongo db running within our docker container
- `> use good_news_db` - so we change (create) to needed db
- `> db.createUser({user: 'suuuper_user', pwd: 'soop3r_U$eR_PSWD', roles:[{role:'dbOwner', db:'good_news_db'}]})` - creating a super user 
- `> db.test_collection.insert({ test: "test" })` - inserting test data to test collection
- `> show collections` - displaying all collections of our previously created db in order to make sure that our test collection was successfully created
- `> exit` - saying goodbye to mongo shell
- `> docker ps` - to show all running docker containers. Find `CONTAINER ID` (first column) of container named `api`.
- `> docker stop <api_contrainer_id>` - to stop api container. We will rerun it in the next step.
- `> docker-compose build && docker-compose up -d` - to rebuild our docker container with new settings applied to MongoDB.
- all credentials for MongoDB should be the same with those which are located in `.env` file in the root of the project (in case if you would like to change them)
- so now you can open a browser window and go to `localhost:6969/v1/news/sources`. Now you should see news sources which are pre-filled before server launch. After 3 minutes, you will be able to see first news gathered from parsing news sites by opening `localhost:6969/v1/news`. Why 3 minutes? Follow my articles on Medium to understand that 😉

# Contact me
If you have any comments or suggestions, please feel free to email me at [batr@ggc.team](mailto:batr@ggc.team) 🙂
If you would like to know when I post new articles, follow me on [twitter 🐦](https://twitter.com/kanzitelli)