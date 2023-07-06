BEGIN;

alter table students
    add create_at timestamp with time zone default now();

alter table students
    add update_at timestamp with time zone;

alter table accounts
    add create_at timestamp with time zone default now();

alter table accounts
    add update_at timestamp with time zone;

alter table credit_limits
    add create_at timestamp with time zone default now();

alter table credit_limits
    add update_at timestamp with time zone;

alter table tasks
    add create_at timestamp with time zone default now();

alter table tasks
    add update_at timestamp with time zone;

COMMIT;