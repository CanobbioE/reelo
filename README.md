# Reelo

Reelo is an ELO system for Mathematics games. The name refers to the Esperanto's term that means "real number".

## Tech stack

### Front end

The UI uses ReactJS with Redux and Hooks. It follows the material UI standards.

### Back end

The logic behind the application is written in [Go](https://golang.org/). Data is being stored in a MySQL database. It follows the REST architecture.

### Devops

For an easy deploy the application is containerized with Docker, and everything is put together with docker-compose.

## Usage

The web-app is deployed on a test server. It is currently in alpha.
Although you can deploy the project on your machine.

## Deploy

To easily test the web application on your machine you need [Docker](https://www.docker.com/get-started) and [Docker-compose](https://github.com/docker/compose).
The [deploy folder](./deploy) contains two bash scripts to start the application and the docker-compose YML file.

### First start

When starting the web-app for the first time, if you are on a unix based machine, use:

```
$ ./deploy-clean-db.sh
```

(Note that this script clears all the images, meaning all the data saved in the database will be lost)
Otherwise you can start the composed docker with:

```
docker-compose up
```

### Keeping an existing database

If you want to keep the data stored in a previously created image of the database, you can restart the application with:

```
$ ./deploy.sh
```

### Configuration files

Inside the [deploy folder](./deploy) you can find three configuration files:

- [.env](./deploy/.env): it's used to configure OS variables for docker. Check the YML file to see which flags are configurable.
- [default.conf](./deploy/default.conf): it's used for ngix server configuration. I recommend to not change that.
- [init.sql](./deploy/init.sql): it's used to initialize the database. I recommend to change the [password's hash](https://github.com/CanobbioE/reelo/blob/8afde13914ef70db072e086907e376350fe39a53/deploy/init.sql#L87). If you don't have an hashing tool, you can find one in the [hashUtils folder](./deploy/hashUtils).

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
