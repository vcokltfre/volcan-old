# Docker Compose Setup

## 1. Clone the bot respository

Navigate to the directory you wish to clone the bot to then run:

```bash
git clone https://github.com/vcokltfre/volcan.git
```

Next, copy the `.env.example` file ready to fill out in later steps:

```bash
cp .env.example .env
```

## 2. Acquire credentials from Discord for the bot and OAuth

Head to the [Discord Developer Portal](https://discord.com/developers/applications) and either create a new application or use one that already exists. You'll need the following data from the Bot and OAuth2 pages:

- Token
- Client ID
- Client Secret

These should be filled into the relevant fields in the `.env` file you created earlier.

## 3. Configure the bot

Copy the `config.example.yml` file and make the relevant changes to configure the bot how you'd like it:

```bash
cp config.example.yml config.yml
```

A config file reference can be found [here](/config).

## 4. Run the bot

You will first need to install docker and docker-compose:

```bash
sudo apt install docker.io docker-compose
```

You can now start the bot by running:

```bash
sudo docker-compose up --build -d
```

!!! Note

    The `-d` flag runs the containers in the background, meaning you don't need to keep the terminal open.
