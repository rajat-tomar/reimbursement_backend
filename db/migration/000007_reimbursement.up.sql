CREATE TABLE reimbursements
(
    id           SERIAL PRIMARY KEY,
    amount       INTEGER,
    user_id      INTEGER,
    expense_id   INTEGER,
    status       TEXT,
    processed_on DATE,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (expense_id) REFERENCES expenses (id)
);