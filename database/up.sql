CREATE TABLE "accounts"
(
  "id"   CHAR(27)    NOT NULL,
  "name" VARCHAR(64) NOT NULL,

  PRIMARY KEY ("id")
);

CREATE TABLE "products"
(
  "id"         CHAR(27)                 NOT NULL,
  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "name"       VARCHAR(64)              NOT NULL,
  "price"      MONEY                    NOT NULL,

  PRIMARY KEY ("id")
);
