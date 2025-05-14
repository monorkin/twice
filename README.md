# Twice

Twice is a clone of the distribution system used by [Once](https://once.com) that's used with 
[Campfire](https://once.com/campfire) 
and [Writebook](https://once.com/writebook)- hence the name, this is a second implementation of Once.

## Repo organization
This is a mono-repo that contains 3 projects:
1. Auth
2. CLI
3. Registry

**Auth** handles 3 tasks - it provides installation scripts and binaries, it is the 
auth backend for the Registry, and it manages products and license keys.

**CLI** installs a product associated with a given license key.

**Registry** is a Docker Registry that hosts app images for products.

## Setup

**First**, from the root of the project run `bin/setup`.

This will do a couple of things needed to get the Registry to talk to Auth.

**Then**, in a separate terminal window run `docker compose up` to start the registry.

Check out the docker-compose.yml file to see how the registry is configured, and 
change it for your OS.

**Finally**, in another terminal window enter the auth directory and 
run:

```bash
bundle
bin/rails db:create db:migrate db:fixtures:load
bin/dev
```

This will install all the necessary depenencies, create the database, run the migrations,
and then start the server on [localhost:3000](http://localhost:3000).

## Usage

You can login with any Developer account to the Auth UI.
Out-of-the-box there are several accounts already pre-populated
(if you've run bin/rails `db:fixtures:load`). You can take a look
at [developers.yml](auth/db/fixtures/developers.yml) to see which
accounts exists and what their credentials are.

In the UI you can create products, developers, customers, and
purchase licenses for customers.

Using the same credentials you can login to the Registry and
push any Docker image you like to it.

```bash
docker login --username jezrien@example.com --password password http://localhost:5000

docker pull hello-world
docker tag hello-world http://localhost:5000/hello-world

docker push http://localhost:5000/hello-world
```

To install a product using the CLI you have to run the following command from the cli directory:

```bash
go run cmd/twice/main.go setup <license-key>
# e.g. `go run cmd/twice/main.go setup nx1w-qg52-e5t3-kpa4`
```

This will pull the product's Docker image from the registry and spin up a contianer for it

## Details

A detailed explanation of how all of this works can be found [here](https://stanko.io/articles/foo-bar)
