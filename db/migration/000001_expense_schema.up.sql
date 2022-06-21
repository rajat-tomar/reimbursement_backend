create table if not exists expenses
(
    Id         serial primary key,
    Amount     int         not null,
    Created_at timestamptz not null default now()
);