create table dailytasks (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        task_name varchar not null,
        created_at timestamp not null default CURRENT_TIMESTAMP,
        completed_today bool default false
);