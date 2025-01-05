create table if not exists contacts
(
    id         serial primary key,
    name       varchar(255) not null unique,
    phone      varchar(20),
    telegram   varchar(50),
    birthday   date,
    created_at timestamptz  not null default now(),
    updated_at timestamptz  not null default now()
);
