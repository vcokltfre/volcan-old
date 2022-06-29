# Manual Setup

## 1. Clone the bot respository

Navigate to the directory you wish to clone the bot to then run:

```bash
git clone https://github.com/vcokltfre/volcan.git
```

Next, copy the `.env.example` file ready to fill out in later steps:

```bash
cp .env.example .env
```

## 2. Set up a Postgres or MySQL (or MariaDB) database

- For Postgres see [this DigitalOcean tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-Ubuntu-20-04).
- For MySQL see [this DigitalOcean tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-Ubuntu-20-04).
- For MariaDB see [this DigitalOcean tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-mariadb-on-Ubuntu-20-04).

Make sure to put the URI of the database you set up in your `.env` file in the `DB_DSN` key.

## 3. Acquire credentials from Discord for the bot and OAuth

Head to the [Discord Developer Portal](https://discord.com/developers/applications) and either create a new application or use one that already exists. You'll need the following data from the Bot and OAuth2 pages:

- Token
- Client ID
- Client Secret

These should be filled into the relevant fields in the `.env` file you created earlier.

## 4. Configure the bot

Copy the `config.example.yml` file and make the relevant changes to configure the bot how you'd like it:

```bash
cp config.example.yml config.yml
```

A config file reference can be found [here](/config).

## 5. Run the bot

!!! Note

    You will need Go installed to build the bot outside Docker. You can see installation instructions for Ubuntu [here](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04).

The bot can now be built and run:

```bash
go build
./volcan
```
