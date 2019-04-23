CREATE TABLE subscriptions (
	user_id UUID NOT NULL,
	repo_id UUID NOT NULL,
	CONSTRAINT "subscriptions_repo_id_fkey" FOREIGN KEY ("repo_id") REFERENCES "public"."repositories"("id"),
    CONSTRAINT "subscriptions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users"("id"),
	PRIMARY KEY (user_id, repo_id)
);
