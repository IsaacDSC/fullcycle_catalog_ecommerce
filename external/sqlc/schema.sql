CREATE TABLE IF NOT EXISTS "categories" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "active" BOOLEAN DEFAULT true,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "products" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "code" VARCHAR(255),
  "name" VARCHAR(255) NOT NULL,
  "image_url" VARCHAR(255),
  "price" INTEGER NOT NULL,
  "description" VARCHAR(255),
  "active" BOOLEAN DEFAULT true,
  "category_id" VARCHAR(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);