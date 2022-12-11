create function public.getdate() returns timestamptz
       stable language sql as 'select now()';

CREATE TABLE "feature" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "active" BOOLEAN NOT NULL,
  "responsible" INT NOT NULL REFERENCES users(id),
  "creation_date" DATE NOT NULL default getdate(),
  "update_date" DATE NULL default getdate()
);