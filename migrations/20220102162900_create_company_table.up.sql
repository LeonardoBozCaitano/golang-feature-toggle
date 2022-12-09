CREATE TABLE "company" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL REFERENCES users(id),
  "name" TEXT NOT NULL,
  "cnpj" TEXT NOT NULL,
  "phone" TEXT NOT NULL,
  "city" TEXT NOT NULL,
  "address" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "image" TEXT NOT NULL,
  "facebook_url" TEXT NOT NULL,
  "instagram_url" TEXT NOT NULL,
  "website_url" TEXT NOT NULL
);