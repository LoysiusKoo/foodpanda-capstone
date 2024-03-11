CREATE TABLE "restaurants" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "address" varchar,
  "rating" varchar,
  "cuisine" varchar,
  "image_url" varchar
);

CREATE TABLE "dishes" (
  "id" bigserial PRIMARY KEY,
  "restaurant_id" bigserial NOT NULL,
  "is_available" bool NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "price" decimal(10, 2) NOT NULL,
  "cuisine" varchar,
  "image_url" varchar
);

CREATE TABLE "cuisine" (
  "id" bigserial PRIMARY KEY,
  "cuisine" varchar,
);


COMMENT ON TABLE "restaurants" IS 'Stores restaurants';

COMMENT ON TABLE "dishes" IS 'Stores dishes';

COMMENT ON TABLE "cuisine" IS 'Stores cuisine';


ALTER TABLE "dishes" ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurants" ("id");

ALTER TABLE "cuisine" ADD FOREIGN KEY ("cuisine") REFERENCES "dishes" ("cuisine");