# encore-ent-template

A starter template combining the [Encore](https://encore.dev) backend framework with the [ent](https://entgo.io) ORM and [Atlas](https://atlasgo.io) versioned migrations.

**Stack:**
- **Encore** – service framework, automatic DB provisioning, local dev server
- **ent** – type-safe Go ORM with code generation
- **Atlas** – schema-diff migration generation (golang-migrate format)
- **PostgreSQL** – via Encore's managed local DB

---

## Prerequisites

| Tool | Install |
|------|---------|
| [Go 1.24+](https://go.dev/dl/) | `brew install go` |
| [Docker](https://docs.docker.com/desktop/install/mac-install/) | Required by Encore for local Postgres |
| [Encore CLI](https://encore.dev/docs/install) | `brew install encoredev/tap/encore` |
| [Atlas CLI v1.x](https://atlasgo.io/getting-started) | `curl -sSf https://atlasgo.sh \| sh` |
| [Task](https://taskfile.dev/installation/) | `brew install go-task` |

> **Keep tools in sync:** after updating the Encore CLI (`encore version update`), also run  
> `go get encore.dev@<new-version> && go mod tidy` so the SDK in `go.mod` matches your CLI.

---

## Getting started with this template

### 1. Create your Encore app

If you want Encore Cloud features (deployment, metrics, secrets), register the app:

```bash
encore app create
```

This writes an app ID into `encore.app`. For local-only development you can skip this step — Encore works fully offline without an ID.

### 2. Install Go dependencies

```bash
go mod download
```

### 3. Start the server

```bash
task start
```

Encore will spin up a local Postgres container, apply the migrations in `api/migrations/`, and start the development server. The local dashboard is available at **http://localhost:9400**.

---

## Task reference

Scripts are managed via [Taskfile](https://taskfile.dev). Run `task --list` for a summary.

| Command | Description |
|---------|-------------|
| `task start` | Start the Encore dev server |
| `task gen` | Re-run all code generators (`go generate ./...`) |
| `task gen-ent` | Regenerate the ent client from schema changes |
| `task gen-new-schema -- Name` | Scaffold a new ent schema entity |
| `task gen-migration -- name` | Generate an Atlas migration from schema diff |
| `task db-reset` | Drop & recreate all databases (dev only) |
| `task db-shell` | Open a psql shell on the `api` database |
| `task tidy` | Sync `go.mod` / `go.sum` |
| `task update-deps` | Upgrade all direct deps to latest |

---

## Typical development workflow

### Add a new entity

```bash
# 1. Scaffold the schema
task gen-new-schema -- Post

# 2. Edit api/ent/schema/post.go to define fields/edges

# 3. Regenerate the ent client
task gen-ent

# 4. Generate a migration
task gen-migration -- add_posts
```

> If you get a "db cluster not running" error during migration generation, restart the
> shadow DB with: `encore db conn-uri --shadow api`

### Apply migrations

Migrations are applied **automatically on startup** by Encore. Just run:

```bash
task start
```

---

## Project layout

```
.
├── encore.app              # Encore app config (set "id" after running encore app create)
├── go.mod / go.sum         # Go module (includes tool deps via tools.go)
├── tools.go                # Pins code-gen tool versions (ent, atlas-provider-ent)
├── Taskfile.yml            # Developer task runner
└── api/
    ├── api.go              # Encore service definition + DB connection
    ├── atlas.hcl           # Atlas migration config (local env)
    ├── migrations/         # golang-migrate SQL files (auto-applied by Encore on boot)
    ├── scripts/
    │   └── generate-migration.sh
    └── ent/
        ├── generate.go     # //go:generate directive for ent
        └── schema/         # ent schema definitions (edit these)
```

---

## Upgrading dependencies

```bash
# Upgrade all direct deps
task update-deps

# Upgrade Encore SDK to match your CLI
encore version update
go get encore.dev@$(encore version | awk '{print $3}' | tr -d 'v') && go mod tidy

# Upgrade Atlas CLI
curl -sSf https://atlasgo.sh | sh
```
