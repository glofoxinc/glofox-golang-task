-- Suggested schema if you take the SQLite path.
--
-- It is only a starting point. Feel free to change column types, add
-- indexes, or model the data differently — we're more interested in the
-- reasoning behind your choices than in matching this verbatim.

CREATE TABLE IF NOT EXISTS classes (
    id          TEXT    PRIMARY KEY,
    name        TEXT    NOT NULL,
    start_date  TEXT    NOT NULL, -- ISO-8601 date, e.g. "2026-12-01"
    end_date    TEXT    NOT NULL,
    capacity    INTEGER NOT NULL CHECK (capacity > 0)
);

CREATE TABLE IF NOT EXISTS bookings (
    id           TEXT    PRIMARY KEY,
    class_id     TEXT    NOT NULL REFERENCES classes(id),
    member_name  TEXT    NOT NULL,
    date         TEXT    NOT NULL, -- ISO-8601 date
    created_at   TEXT    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS bookings_by_class_date
    ON bookings (class_id, date);
