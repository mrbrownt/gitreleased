CREATE TABLE repositories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	owner VARCHAR NULL,
	name VARCHAR NULL,
	github_id VARCHAR NULL,
	description VARCHAR NULL,
	url VARCHAR NULL
);
