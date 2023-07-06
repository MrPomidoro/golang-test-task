BEGIN;

CREATE INDEX idx_accounts_student_id ON accounts(student_id);
CREATE INDEX idx_credit_limits_student_id ON credit_limits(student_id);
CREATE INDEX idx_tasks_id ON tasks(id);

COMMIT;