CREATE SEQUENCE IF NOT EXISTS releases_id_seq;

CREATE TABLE releases (
	"id" int4 PRIMARY KEY DEFAULT nextval('releases_id_seq'::regclass),
	name VARCHAR NOT NULL,
	major INT NOT NULL DEFAULT '0',
	minor INT NOT NULL DEFAULT '0',
	patch INT NOT NULL DEFAULT '0',
	dev VARCHAR NULL
);

CREATE TABLE release_to_repo (
	release_id INT NOT NULL,
	repo_id UUID NOT NULL,
	CONSTRAINT "release_to_repo_release_id_fkey" FOREIGN KEY ("release_id") REFERENCES "public"."releases"("id"),
	CONSTRAINT "release_to_repo_repo_id_fkey" FOREIGN KEY ("repo_id") REFERENCES "public"."repositories"("id"),
	PRIMARY KEY (release_id, repo_id)
);
