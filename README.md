# Dev Machine 

# Instruction

- First step you should create your machine that will host our containers  

```sh
$ docker-machine create --driver virtualbox dev
```

- Configure the terminal to use docker dev machine 

```sh
$ docker-machine env dev
```

- Connect your shell to the dev machine.

```sh
$ eval $(docker-machine env dev)
```

