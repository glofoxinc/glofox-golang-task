# Glofox — Backend Live Coding Task

Welcome, and thanks for taking the time to interview with us.

This is a **30-minute live coding exercise** in Go. We'll work through it together — coding, talking, and iterating. The goal isn't to finish a checklist; it's to give us a window into how you think and how you write Go.

## The domain

Glofox is a SaaS platform for boutiques, studios, and gyms. Studio owners create classes; members book onto specific dates.

## Your task

Build a small HTTP API with one endpoint.

### Pre-seeded class

A Pilates class is already loaded into the store when the server starts — you don't need to create it. Its details are:

| Field | Value |
|-------|-------|
| `id`          | `5d2e2a9f-3c1b-4f2a-b6e0-1a2b3c4d5e6f` |
| `name`        | `Pilates`                              |
| `start_date` | `2026-12-14`                           |
| `end_date`   | `2026-12-14`                           |
| `capacity`    | `10`                                   |

The class is a single instance on `2026-12-14` with a maximum of 10 attendees. The ID is also printed to stdout when you run the server.

### `POST /bookings` — book a member onto the class

A member books a specific date of the pre-seeded class.

```http
POST /bookings
Content-Type: application/json

{
  "member_name": "Alice",
  "class_id":    "5d2e2a9f-3c1b-4f2a-b6e0-1a2b3c4d5e6f",
  "date":        "2026-12-14"
}
```

Respond with an appropriate success status. Think about what should happen when the inputs don't make sense — both shape-wise and meaning-wise.

## Storage

Use **in-memory** storage. [`internal/storage/memory/`](internal/storage/memory/) is the empty package waiting for it.

## What we value

The task above is the floor, not the ceiling. As you work, we'll be paying attention to:

- **Correctness under contention.** What happens when two members try to book the last seat at the same instant?
- **Where the rules live.** Is "this booking is valid" enforced in the handler, the service, or the storage layer? Why?
- **Persistence boundaries.** Where does your domain end and storage begin? How hard would it be to swap in a SQL-backed store later?
- **How you test.** You won't have time to test everything. Pick one test that you think matters — and we'll ask why you picked it.
- **Trade-offs you can defend.** We'd rather see a simple solution you can justify than a clever one you can't.

You don't need to solve all of these upfront. We'll explore some of them together once the basic endpoint is working.

## A note on AI tools

You're welcome to use AI assistants (Cursor, Copilot, ChatGPT, anything else). We'll ask you to walk through your code and the choices behind it — so use AI the way you'd use it on the job: as a collaborator whose output you understand and own.

If you'd rather not use AI, that's also completely fine.

## Running the skeleton

Requires **Go 1.23+**.

```bash
make run         # starts the server on :8080
make test        # runs your tests once you write them
make test-race   # runs tests with the race detector
```

No tests ship in this repo — write the ones you think matter.

## Layout

```text
.
├── cmd/api/main.go              # boots an http.Server on :8080; holds the seed class data
├── internal/
│   └── storage/
│       └── memory/              # empty — fill in your in-memory storage here
├── go.mod                       # stdlib only
└── Makefile
```

The `storage/memory` package is deliberately empty. We'd like to see how you'd organise the domain, service, and HTTP layers around it.

The skeleton uses the Go standard library only. You're welcome to reach for `chi`, `gin`, `echo`, or anything else — we'll ask why.

## Submitting

You don't need to. We'll be coding together; the conversation matters more than what's on disk at the end. Push to a branch on this repo (or a fork) if you want a record of where we got to.

Good luck, and have fun.
