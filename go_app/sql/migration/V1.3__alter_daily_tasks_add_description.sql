CREATE TABLE IF NOT EXISTS dailytaskrunbook (
    id                   VARCHAR(100) PRIMARY KEY NOT NULL,
    runbook_task              TEXT NOT NULL,
    created_at           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    task_id              VARCHAR(100),
    FOREIGN KEY (task_id) REFERENCES dailytasks(id)
    );