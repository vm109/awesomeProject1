CREATE TABLE plans(
    id SERIAL PRIMARY KEY,
    plan_description jsonb,
    created_at timestamp,
    updated_at timestamp,
    progress_phase text
)