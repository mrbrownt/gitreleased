CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	email STRING NULL,
	github_id STRING NULL,
	github_user_name STRING NULL,
	first_name STRING NULL,
	last_name STRING NULL,
	access_token STRING NULL,
    UNIQUE INDEX github_user_name_idx (github_user_name ASC)
);
