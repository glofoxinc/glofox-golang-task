# Glofox — Backend Live Coding Task

Welcome, and thanks for taking the time to interview with us.

This is a **30-minute live coding exercise** in Go. We'll work through it together — coding, talking, and iterating. The goal isn't to finish a checklist; it's to give us a window into how you think and how you write Go.

## The domain

Glofox is a SaaS platform for boutiques, studios, and gyms. Studio owners create classes; members book onto specific dates.

## Your task

Build a small HTTP API with two endpoints.

### `POST /classes` — create a class

A studio owner creates a class with a date range and a per-day capacity.

```http
POST /classes
Content-Type: application/json

{
  "name": "Pilates",
  "start_date": "2026-12-01",
  "end_date":   "2026-12-20",
  "capacity":   10
}
```

A class with `start_date = 2026-12-01`, `end_date = 2026-12-20`, `capacity = 10` means there are 20 daily class instances, each with a maximum of 10 attendees.

Respond with an appropriate success status and an identifier the client can refer to later.

### `POST /bookings` — book a member onto a class

A member books a specific date of an existing class.

```http
POST /bookings
Content-Type: application/json

{
  "member_name": "Alice",
  "class_id":    "<id from POST /classes>",
  "date":        "2026-12-14"
}
```

Respond with an appropriate success status. Think about what should happen when the inputs don't make sense — both shape-wise and meaning-wise.

## Storage

Start in-memory. There is no database to set up. If you want to talk through what changes when this is backed by a real store, we'd love to hear it.

## What we value

The task above is the floor, not the ceiling. As you work, we'll be paying attention to:

- **Correctness under contention.** What happens when two members try to book the last seat at the same instant?
- **Where the rules live.** Is "this booking is valid" enforced in the handler, the service, or the storage layer? Why?
- **Persistence boundaries.** Your code is in-memory today. How hard would it be to swap in a SQL or document store tomorrow?
- **Trade-offs you can defend.** We'd rather see a simple solution you can justify than a clever one you can't.

You don't need to solve all of these upfront. We'll explore some of them together once the basic endpoints are working.

## A note on AI tools

You're welcome to use AI assistants (Cursor, Copilot, ChatGPT, anything else). We'll ask you to walk through your code and the choices behind it — so use AI the way you'd use it on the job: as a collaborator whose output you understand and own.

If you'd rather not use AI, that's also completely fine.

## Running the skeleton

Requires **Go 1.22+**.

```bash
make run         # starts the server on :8080
make test        # runs the smoke test (it FAILS until /classes is implemented)
make test-race   # runs tests with the race detector
```

The repo ships with a single failing black-box test in [`api_test.go`](api_test.go) that POSTs to `/classes` and expects `201 Created`. Make it pass first; that's the warm-up.

## Layout

```text
.
├── cmd/api/main.go   # boots an http.Server on :8080 with no routes
├── internal/         # empty — your call on package layout
├── api_test.go       # one failing smoke test
├── go.mod            # stdlib only; add deps if you want them
└── Makefile
```

`internal/` is empty deliberately. We'd like to see how you'd organise this.

The skeleton uses only the Go standard library. You're welcome to reach for `chi`, `gin`, `echo`, or anything else — we'll ask why.

## Submitting

You don't need to. We'll be coding together; the conversation matters more than what's on disk at the end. Push to a branch on this repo (or a fork) if you want a record of where we got to.

Good luck, and have fun.
