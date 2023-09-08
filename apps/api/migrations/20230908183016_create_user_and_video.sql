-- Create "users" table
CREATE TABLE "users" ("user_id" uuid NOT NULL, "username" character varying NOT NULL, "password" character varying NOT NULL, "name" character varying NOT NULL, "registered_at" timestamptz NOT NULL, PRIMARY KEY ("user_id"));
-- Create index "user_username_password" to table: "users"
CREATE INDEX "user_username_password" ON "users" ("username", "password");
-- Create index "users_username_key" to table: "users"
CREATE UNIQUE INDEX "users_username_key" ON "users" ("username");
-- Create "videos" table
CREATE TABLE "videos" ("video_id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "video_title" character varying NOT NULL, "video_url" character varying NOT NULL, "video_thumbnail" character varying NULL, "creator_name" character varying NOT NULL, "creator_url" character varying NOT NULL, "created_at" timestamptz NOT NULL, "posted_by" uuid NOT NULL, PRIMARY KEY ("video_id"), CONSTRAINT "videos_users_videos" FOREIGN KEY ("posted_by") REFERENCES "users" ("user_id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "video_video_url" to table: "videos"
CREATE INDEX "video_video_url" ON "videos" ("video_url");
-- Create index "videos_video_url_key" to table: "videos"
CREATE UNIQUE INDEX "videos_video_url_key" ON "videos" ("video_url");
