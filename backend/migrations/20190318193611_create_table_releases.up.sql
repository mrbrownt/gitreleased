CREATE TABLE releases (
	id INT NOT NULL DEFAULT unique_rowid(),
	name STRING NOT NULL,
	major INT NOT NULL DEFAULT '0',
	minor INT NOT NULL DEFAULT '0',
	patch INT NOT NULL DEFAULT '0',
	dev STRING NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, name, major, minor, patch, dev)
);

CREATE TABLE release_to_repo (
	id INT NOT NULL DEFAULT unique_rowid(),
	release_id INT NOT NULL,
	repo_id UUID NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	CONSTRAINT fk_release_id_ref_releases FOREIGN KEY (release_id) REFERENCES releases (id) ON DELETE CASCADE,
	UNIQUE INDEX release_idx (release_id ASC),
	CONSTRAINT fk_repo_id_ref_repositories FOREIGN KEY (repo_id) REFERENCES repositories (id),
	INDEX repo_idx (repo_id ASC),
	FAMILY "primary" (id, release_id, repo_id)
);
