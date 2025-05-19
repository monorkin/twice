# Twice

Twice is a clone of the [Once](https://once.com) distribution system 
used by [Campfire](https://once.com/campfire) and [Writebook](https://once.com/writebook).

Through Twice you can generate and distribute license keys with which
users can install and run Docker containers of your products on their servers.

## Example

Twice gives you the ability to issue license keys
that can be used to install a product.

Products are Docker containers installed from Docker images hosted on a 
private Docker Registry.

The whole installation process looks like this:

```bash
# Download and install the Twice CLI,
# and then install a product with the license key nx1w-qg52-e5t3-kpa4 from auth.example.com
/bin/bash -c "$(curl -fsSL http://auth.example.com/install/nx1w-qg52-e5t3-kpa4)"

# Install a product with the license key nx1w-qg52-e5t3-kpa4 from example.org
twice setup nx1w-qg52-e5t3-kpa4@example.org
```

You can distribute one or more products this way, and even copies of the same product.

## About

A detailed explanation of how all of this works, and how Twice came to be, can be found [here](https://stanko.io/building-twice-a-clone-of-once-gJKxLYCe26Ak).
[Release 0.1.0](https://github.com/monorkin/twice/releases/tag/0.1.0) contains all the code referenced in that article.

## Repo organization
This is a mono-repo that contains 3 projects:
1. Auth
2. CLI
3. Registry

**Auth** handles 3 tasks - it provides installation scripts and binaries, it is the 
auth backend for the Registry, and it manages products and license keys.

**CLI** installs a product associated with a given license key and manages all installed products.

**Registry** is a Docker Registry that hosts app images for products.

## Development

**First**, make sure you have the following installed:
1. Ruby 3.4 or newer
2. Go 1.23 or newer
3. Docker
4. Docker Compose

**Second**, from the root of the project run `bin/setup`.

This will setup the Registry and the Auth server so that they can talk to one another.
And it will compile the CLI, and prepare it to be distributed using the Auth server.

**Third**, in a separate terminal window, from the `registry` directory,
run `docker compose up` to start the registry.

Check out the docker-compose.yml file to see how the registry is configured, and 
change it for your OS.

**Fourth**, in another terminal window from the `auth` directory run:

```bash
bundle
bin/rails db:create db:migrate db:fixtures:load
bin/dev
```

This will install all the necessary dependencies, create the database, run the migrations,
and then start the server on [localhost:3000](http://localhost:3000).

## Usage

### Auth
You can login with any Developer account to the Auth UI.

Out-of-the-box there are several accounts already pre-populated
(if you've run bin/rails `db:fixtures:load`). You can take a look
at [developers.yml](auth/db/fixtures/developers.yml) to see which
accounts exists and what their credentials are.

In the UI you can create products, developers, customers, and
purchase licenses for customers.

You can also install a product directly from the auth server using this command:

```bash
/bin/bash -c "$(curl -fsSL http://localhost:3000/install/nx1w-qg52-e5t3-kpa4)"
```

This will download and install the cli binary, and then run it, passing it
the license key from the URL.

### Registry

Using the same credentials you can login to the Registry and
push any Docker image you like to it.

Just note that only developer accounts have push access to the registry

```bash
docker login --username jezrien@example.com --password password http://localhost:5000

docker pull hello-world
docker tag hello-world http://localhost:5000/hello-world

docker push http://localhost:5000/hello-world
```

To delete images from the registry you can use

```bash
bin/delete-image localhost:5000 hello-world jezrien@example.com password
```

### CLI

To install a product using the CLI you have to run the following command from the `cli` directory:

```bash
go run cmd/twice/main.go setup <license-key>
# e.g. `go run cmd/twice/main.go setup nx1w-qg52-e5t3-kpa4`
```

This will pull the product's Docker image from the registry and spin up a container for it
## Uninstall

To uninstall the CLI run

```bash
rm /usr/local/bin/twice
```

That's it.

