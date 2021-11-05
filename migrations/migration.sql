DROP TABLE IF EXISTS authors CASCADE;
DROP TABLE IF EXISTS posts;

CREATE TABLE authors (
    "id"     SERIAL,
    "name"   TEXT NOT NULL,
    "email"  TEXT NOT NULL,
    -- "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- "deleted_at" TIMESTAMP(3)

    CONSTRAINT "authors_pkey" PRIMARY KEY ("id")
);

CREATE TABLE posts (
    "id"        SERIAL,
    "title"     TEXT NOT NULL,
    "text"      TEXT NOT NULL,
    "author_id" SERIAL NOT NULL,
    -- "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- "deleted_at" TIMESTAMP(3)

    CONSTRAINT "posts_pkey" PRIMARY KEY ("id")
);

-- unique email
CREATE UNIQUE INDEX "name_email_key" ON "authors"("email");

-- unique name
CREATE UNIQUE INDEX "name_authors_key" ON "authors"("name");

ALTER TABLE "posts" ADD CONSTRAINT "posts_authorid_fkey" FOREIGN KEY ("author_id") REFERENCES "authors"("id") ON DELETE RESTRICT ON UPDATE CASCADE;