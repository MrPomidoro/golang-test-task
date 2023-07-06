BEGIN;

CREATE TABLE students (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  full_name TEXT NOT NULL,
  group_num TEXT NOT NULL,
  email TEXT NOT NULL,
  username TEXT NOT NULL UNIQUE
);

COMMIT;