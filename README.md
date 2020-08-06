# Reelo

Reelo is an ELO system for Mathematics games. The name refers to the Esperanto's term that means "real number".

# NOTE: This project has been refactored into different private repositories. This repo will be kept as a reference to the first prototype.

## Tech stack
### Front end
The UI uses ReactJS with Redux and Hooks. It follows the material UI standards.

### Back end
The logic behind the application is written in [Go](https://golang.org/). Data is being stored in a MySQL database. It follows the REST architecture.

### Deploy
The application is containerized with Docker, and everything is put together with docker-compose.

## Usage
The web-app is deployed on a test server. It is currently in alpha.
Although you can deploy the project on your machine.

## Deploy
To easily test the web application on your machine you need [Docker](https://www.docker.com/get-started) and [Docker-compose](https://github.com/docker/compose).
The [deploy folder](./deploy) contains two bash scripts that use the images stored on DockerHub to start the aplication. In the same older, there is also the docker-compose YML file.

### Start with DockerHub images
When starting the web-app for the first time, if you are on a unix based machine, use:
```bash
$ ./deploy-clean-db.sh
```

This script will first clear all the docker images/containers related to this project, then it will pull the new images form DockerHub.
This means that all the data saved within the database image will be lost. You can start the composed docker with `$ ./deploy.sh` to preserve data stored within an already existing database image.

**Note**: Since the script makes use of images stored in the DockerHub registry, you cannot apply any change you make locally.

### Configuration files
Inside the [deploy folder](./deploy) you can find three configuration files:

- [.env](./deploy/.env): it's used to configure OS variables for docker. Check the YML file to see which flags are configurable.
- [default.conf](./deploy/default.conf): it's used for ngix server configuration. I recommend to not change that.
- [init.sql](./deploy/init.sql): it's used to initialize the database. I recommend to change the [password's hash](https://github.com/CanobbioE/reelo/blob/8afde13914ef70db072e086907e376350fe39a53/deploy/init.sql#L87). If you don't have an hashing tool, you can find one in the [hashUtils folder](./deploy/hashUtils).

### Build and run from source for testing
If you want to build from source you'll still need [Docker](https://www.docker.com/get-started) and [Docker-compose](https://github.com/docker/compose), as well as [Go](https://golang.org/) and [npm](https://www.npmjs.com/).
First of all start the database from the `deploy` folder:
```bash
$ docker-compose up -d reelo-db
```
Inside the `backend` folder run:
```bash
$ go build
$ ./backend
```
Inside the `frontend/reelo folder`, if you'd like to have front end hot reloading, run:
```bash
$ npm install
$ npm run
```
otherwise (no hot reloading) just run
```bash
$ npm build
```
The artifacts generated should be copied inside your local http server folder.

Visiting <http://localhost/> will display the web-application.

API calls can be manually sent to <http://localhost:8080>.




## TODO

### Back end

- API limiter (?)
- Better auth handling
- Paris

### Front end

- i18n (?)
- Caching (~)
- Better cookies handling
- A bit of refactoring wouldn't hurt
- Ranks filters (order by, search, show only category, show year)
- TypeScript (?)

## Credits

- Ideation: Cesco Reale
- Implementation: Fabio Angelini, Anna Bernardi, Edoardo Canobbio
- Scientific Committee: Fabio Angelini, Andrea Nari, Marco Pellegrini, Cesco Reale
- Collaborators: David Barbato, Maurizio De Leo, Francesco Morandin, Alberto Saracco
