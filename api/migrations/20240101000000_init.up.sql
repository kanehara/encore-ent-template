CREATE TABLE IF NOT EXISTS "users" (
  "id"          uuid                        NOT NULL,
  "delete_time" timestamp with time zone    NULL,
  "name"        character varying           NOT NULL,
  PRIMARY KEY ("id")
);
