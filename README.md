# example-go

## Installation & Setup

## API Documentation

- Follow [API Contract](https://hackmd.io/sYb9vrFtRV6l0hiBm_Hvpw) for API Documentation

### Note: Please use the following command before doing any commit

### clone repository

```bash
git clone $REPOSITORY
cd $REPOSITORY
git config core.hooksPath .githooks
chmod ug+x .githooks/pre-commit
chmod ug+x .githooks/commit-msg
```

### download dependency

```bash
go mod vendor -v
```

### run necessary service container

```bash
cp compose.example.yaml compose.yaml
docker-compose up -d
```

## 1. RUN
```bash
air
```
### make build file executable

```bash
chmod +x build.sh
```

### build project

```bash
./build.sh
```

### run migration

```bash
./example-go migration up
```

### run application; note this will ask for a config.yaml file. make a config.yaml file from example and change accordingly

```bash
./example-go serve
```

## 2. OR RUN using makefile

```bash
make migration-up
make build-run
```
