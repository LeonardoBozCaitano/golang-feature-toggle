CREATE TABLE "feature" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "active" BOOLEAN NOT NULL,
  "responsible" INT NOT NULL REFERENCES users(id),
  "creation_date" DATE NOT NULL,
  "update_date" DATE NULL
);