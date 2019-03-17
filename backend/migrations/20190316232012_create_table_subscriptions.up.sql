CREATE TABLE subscriptions (
	id INT NOT NULL DEFAULT unique_rowid(),
	"user" UUID NOT NULL,
	repo UUID NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	CONSTRAINT fk_user_ref_users FOREIGN KEY ("user") REFERENCES users (id),
	INDEX subscriptions_user_idx ("user" ASC),
	CONSTRAINT fk_repo_ref_repositories FOREIGN KEY (repo) REFERENCES repositories (id),
	INDEX subscriptions_repo_idx (repo ASC),
	FAMILY "primary" (id, "user", repo)
);
