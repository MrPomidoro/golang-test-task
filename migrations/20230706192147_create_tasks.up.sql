BEGIN;

CREATE TABLE tasks (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   description TEXT NOT NULL,
   cost INT NOT NULL
);

COMMIT;