CREATE TABLE "blocks" (
  "block_num" bigint PRIMARY KEY,
  "block_hash" varchar NOT NULL UNIQUE,
  "block_time" bigint NOT NULL,
  "parent_hash" varchar NOT NULL
);

CREATE TABLE "transactions" (
    "tx_hash" varchar NOT NULL PRIMARY KEY,
    "block_hash" varchar NOT NULL,
    "block_num" bigint NOT NULL,
    "from" varchar NOT NULL,
    "to" varchar NOT NULL,
    "nonce" bigint NOT NULL,
    "value" bigint NOT NULL,
    "gas" bigint NOT NULL,
    "tx_index" bigint NOT NULL
);

CREATE TABLE "logs" (
    "id" bigserial PRIMARY KEY,
    "address" varchar NOT NULL,
    "topics" varchar[] NOT NULL,
    "block_num" bigint NOT NULL,
    "tx_hash" varchar NOT NULL,
    "block_hash" varchar NOT NULL,
    "log_index" bigint NOT NULL,
    "removed" boolean NOT NULL
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("block_hash") REFERENCES "blocks" ("block_hash");
ALTER TABLE "transactions" ADD FOREIGN KEY ("block_num") REFERENCES "blocks" ("block_num");

ALTER TABLE "logs" ADD FOREIGN KEY ("block_num") REFERENCES "blocks" ("block_num");

ALTER TABLE "logs" ADD FOREIGN KEY ("tx_hash") REFERENCES "transactions" ("tx_hash");

CREATE INDEX ON "blocks" ("block_num");
CREATE INDEX ON "blocks" ("block_hash");

CREATE INDEX ON "transactions" ("block_hash");
CREATE INDEX ON "transactions" ("block_num");
CREATE INDEX ON "transactions" ("tx_hash");
