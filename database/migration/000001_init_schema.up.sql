CREATE TABLE "restaurants" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "address" varchar,
  "rating" float NOT NULL,
  "restaurant_type" varchar NOT NULL,
  "num_of_reviews" int NOT NULL,
  "image_url" varchar
);

CREATE TABLE "dishes" (
  "id" bigserial PRIMARY KEY,
  "restaurant_id" bigserial NOT NULL,
  "is_available" bool NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "price" decimal(10, 2) NOT NULL,
  "diet_type" varchar NOT NULL,
  "cuisine" varchar NOT NULL,
  "image_url" varchar NOT NULL
);

CREATE TABLE "cuisine" (
  "id" bigserial PRIMARY KEY,
  "cuisine" varchar
);

CREATE TABLE "playlists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "image" varchar NOT NULL,
  "food_items" varchar NOT NULL,
  "is_active" bool NOT NULL DEFAULT true,
  "created_on" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "playlist_dishes" (
  "id" bigserial PRIMARY KEY,
  "playlist_id" bigserial NOT NULL,
  "dish_id" bigserial NOT NULL,
  "date_to_be_delivered" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "added_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "searches" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "keyword" varchar
);


COMMENT ON TABLE "restaurants" IS 'Stores restaurants';

COMMENT ON TABLE "dishes" IS 'Stores dishes';

COMMENT ON TABLE "cuisine" IS 'Stores cuisine';

COMMENT ON TABLE "playlists" IS 'Stores user playlist';


ALTER TABLE "dishes" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurants" ("id");