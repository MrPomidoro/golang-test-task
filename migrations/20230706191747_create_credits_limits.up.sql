BEGIN;

CREATE TABLE credit_limits
(
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   student_id UUID NOT NULL REFERENCES students(id),
   credit_limit INT NOT NULL DEFAULT 1000
);

COMMIT;