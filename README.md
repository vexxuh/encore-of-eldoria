# Cosmic Encore

a discord bot generator. Allowing people to build simple bots and deploy them to their discord server with ease!

## Running locally

Before running your application, make sure you have Docker installed and running. It's required to locally run Encore applications with databases.

```bash
encore run
```

## Open the developer dashboard

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to access Encore's [local developer dashboard](https://encore.dev/docs/observability/dev-dash).

## Deployment

Deploy your application to a staging environment in Encore's free development cloud:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

Then head over to the [Cloud Dashboard](https://app.encore.dev) to monitor your deployment and find your production URL.

## Testing

```bash
encore test ./...
```

Public Whiteboard for documentation here: https://app.whiteboard.microsoft.com/me/whiteboards/ef75538d-9a4b-46a4-b2fd-d2b22fb235a0
