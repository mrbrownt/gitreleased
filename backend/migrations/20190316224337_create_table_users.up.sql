CREATE EXTENSION IF NOT EXISTS  pgcrypto;

CREATE TABLE users (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	email VARCHAR NULL,
	github_id VARCHAR NULL,
	github_user_name VARCHAR NULL,
	first_name VARCHAR NULL,
	last_name VARCHAR NULL,
	access_token VARCHAR NULL
);
