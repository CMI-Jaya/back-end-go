-- Tabel Users
CREATE TABLE "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "employee_id" varchar,
  "name" varchar,
  "email" varchar UNIQUE,
  "password" varchar,
  "phone_number" varchar,
  "profile_picture" varchar,
  "role" varchar,
  "status" varchar CHECK (status IN ('active', 'inactive')),
  "remember_token" varchar,
  "email_verified_at" timestamp,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Categories
CREATE TABLE "categories" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Articles
CREATE TABLE "articles" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "category_id" integer,
  "title" varchar,
  "slug" varchar,
  "tags" json,
  "content" text,
  "message" text,
  "thumbnail" varchar,
  "alt_thumbnail" varchar,
  "banner" varchar,
  "alt_banner" varchar,
  "poster" varchar,
  "alt_poster" varchar,
  "link_video" varchar,
  "status" varchar CHECK (status IN ('published', 'draft', 'archived')),
  "meta_title" varchar,
  "meta_description" varchar,
  "author_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Articles Views
CREATE TABLE "articles_views" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "article_id" integer,
  "ip_address" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Comments
CREATE TABLE "comments" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "article_id" integer,
  "username" varchar,
  "email" varchar,
  "comment" varchar,
  "parent_id" integer,
  "status" varchar CHECK (status IN ('approved', 'pending', 'rejected')),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Visits
CREATE TABLE "visits" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "ip_address" varchar,
  "url" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Testimonials
CREATE TABLE "testimonials" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "comment" varchar,
  "photo_profile" varchar,
  "category_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Appointments
CREATE TABLE "appointments" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "phone_number" varchar,
  "email" varchar,
  "date_of_booking" date,
  "time" timestamp,
  "link_meet" varchar,
  "host_id" integer,
  "pdf_file" varchar,
  "img" varchar,
  "status" varchar CHECK (status IN ('confirmed', 'pending', 'cancelled')),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Webinars
CREATE TABLE "webinars" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "link_meet" varchar,
  "host_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Videos
CREATE TABLE "videos" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "title" varchar,
  "description" varchar,
  "link_video" varchar,
  "category_id" integer,
  "meta_title" varchar,
  "meta_description" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Tabel Notifications
CREATE TABLE "notifications" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer,
  "type" varchar,
  "message" text,
  "status" varchar CHECK (status IN ('unread', 'read')),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

-- Relasi Foreign Key
ALTER TABLE "articles" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "articles" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");
ALTER TABLE "articles_views" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("article_id") REFERENCES "articles" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("parent_id") REFERENCES "comments" ("id");
ALTER TABLE "testimonials" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "videos" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "notifications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "appointments" ADD FOREIGN KEY ("host_id") REFERENCES "users" ("id");
ALTER TABLE "webinars" ADD FOREIGN KEY ("host_id") REFERENCES "users" ("id");