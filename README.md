# Encore Of Eldoria

Encore Of Eldoria is a Discord RPG (soon to be Discord MMO) that allows players to travel with friends, fight for glory, and much more! We got our idea from the Encore community location change to Discord. While Encore has an example of a Slack Bot, they do not have one yet for a Discord Bot, so we built this as a nod to the community location change!

# Project Structure

This bot, while simple (for now), highlights the power of Encore's backend API, infrastructure SDK, databasing, and caching tools.

# Application Overview

## discord-bot

This "package" holds all code related to the Discord bot. While we wish to have deployed this service using Encore's Cloud Deployment system, none of us wish to spend the money at the moment on this hackathon project, so please just imagine that we did and ignore the fact that this is set up for local APIs only.

## eldoria

The Encore package is not just a component, but the very foundation of our 'game'. Through Encore's API and databasing tools, we manage the setup and hosting of our stored game logic and facilitate data transfer between our backend game processing and the Discord bot.

The game processing consists of 2 - 4 parts. First, we pass the parameters from Discord into our game engine to handle the state change of the player's character. Then, the game engine will pass out generated text and/or photos for the player's actions. These tell the story for our players (The art and text are all AI generated using NovelAi. We curate the text generation to ensure the players are on some rails and can't just do whatever they want, changing the lore or what characters and NPCs can say or do, for example). Finally, we commit the player character's state changes to a database so that we can ensure their health and other stats are accurate when the player performs another action!

# Deployment (Local Only (for now))

Before running your application, ensure you have installed Docker and the Encore CLI, which are running/working. Docker is required to run Encore applications with databases locally and is also needed for the Discord bot. The Encore CLI runs the Encore-related dependencies.

```
Docker Install link: https://docs.docker.com/engine/install/
```

```
Encore Install link: https://encore.dev/docs/quick-start
```

Run (from your project root directory) either the `eoe-deploy-local.ps1` or the `eoe-deploy-local.sh` script (depending on your operating system) these will take care of deploying your applicaitons locally.

If you need additional help setting up your discord bot on your server please review this guide!

```
Discord Bot Install/Setup: https://www.ionos.com/digitalguide/server/know-how/creating-discord-bot/
```

Lastly, add your env values need to be configured

(TBD setup guide for env values)

Then thats it you are done! Enjoy the bot!
