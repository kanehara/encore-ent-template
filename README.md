# API

## Getting started

1. [Download Go](https://go.dev/dl/).
2. Install [Docker](https://docs.docker.com/desktop/install/mac-install/).
3. Install [Encore](https://encore.dev/docs/install).
4. Install [Atlas](https://atlasgo.io/getting-started)

## Tasks

Scripts are managed in Taskfile.yml (see docs [here](https://taskfile.dev/usage/#forwarding-cli-arguments-to-commands))

Run the app

```
task start
```

Generate new ent schema

```
task gen-new-schema -- some_schema_name
```

Generate ent client from updated schema

```
task gen-ent
```

Generate new migration from updated schema

```
task gen-migration -- migration_name
```
If you get an error with the db cluster not running, run the following to get the shadow DB running
```
encore db conn-uri --shadow api
```

Apply migrations by shutting down the server and starting it back up

```
task start
```
