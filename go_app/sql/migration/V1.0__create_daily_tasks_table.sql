CREATE TABLE IF NOT EXISTS dailytasks
(
    id                   VARCHAR(26) PRIMARY KEY  NOT NULL,
    task_name varchar not null,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    completed_today bool default false
                                       );