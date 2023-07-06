BEGIN;

CREATE TABLE accounts (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  student_id UUID NOT NULL REFERENCES students(id),
  debt INT NOT NULL DEFAULT 0
);

COMMIT;