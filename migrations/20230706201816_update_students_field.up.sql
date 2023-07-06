BEGIN;

alter table students
    alter column verify_email set default false;

COMMIT;