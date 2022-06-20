create table if not exists expenses
(
    id         bigserial primary key,
    amount     bigint      not null,
    created_at timestamptz not null default now()
);