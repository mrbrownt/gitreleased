CREATE TABLE repositories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	owner STRING NULL,
	name STRING NULL,
	github_id STRING NULL,
	description STRING NULL,
	url STRING NULL,
	UNIQUE INDEX github_id_idx (github_id ASC)
);
