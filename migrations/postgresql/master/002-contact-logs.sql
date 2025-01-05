create table if not exists contact_logs
(
    id          serial primary key,
    contact_id  int         not null references contacts (id),
    datetime    timestamptz not null default now(),
    log_massage text        not null default '',
    created_at  timestamptz not null default now(),
    updated_at  timestamptz not null default now()
);
