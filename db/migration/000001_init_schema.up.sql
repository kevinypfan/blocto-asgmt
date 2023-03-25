CREATE TABLE "blocks" (
  "block_id" bigint PRIMARY KEY,
  "block_hash" varchar NOT NULL UNIQUE,
  "block_number" varchar NOT NULL UNIQUE,
  "difficulty" varchar NOT NULL,
  "extra_data" varchar NOT NULL,
  "gas_limit" varchar NOT NULL,
  "gas_used" varchar NOT NULL,
  "logs_bloom" varchar NOT NULL,
  "miner" varchar NOT NULL,
  "mix_hash" varchar NOT NULL,
  "nonce" varchar NOT NULL,
  "parent_hash" varchar NOT NULL,
  "receipts_root" varchar NOT NULL,
  "sha3_uncles" varchar NOT NULL,
  "size" varchar NOT NULL,
  "state_root" varchar NOT NULL,
  "timestamp" varchar NOT NULL,
  "total_difficulty" varchar NOT NULL,
  "transactions_root" varchar NOT NULL
);

CREATE TABLE "transactions" (
    "transaction_hash" varchar NOT NULL PRIMARY KEY,
    "block_hash" varchar NOT NULL,
    "block_number" varchar NOT NULL,
    "from" varchar NOT NULL,
    "gas" varchar NOT NULL,
    "gas_price" varchar NOT NULL,
    "input" varchar NOT NULL,
    "nonce" varchar NOT NULL,
    "to" varchar NOT NULL,
    "transaction_index" varchar NOT NULL UNIQUE,
    "value" varchar NOT NULL,
    "type" varchar NOT NULL,
    "v" varchar NOT NULL,
    "r" varchar NOT NULL,
    "s" varchar NOT NULL
);

CREATE TABLE "logs" (
    "id" bigserial PRIMARY KEY,
    "address" varchar NOT NULL,
    "topics" varchar NOT NULL,
    "data" varchar NOT NULL,
    "block_number" varchar NOT NULL,
    "transaction_hash" varchar NOT NULL,
    "transaction_index" varchar NOT NULL,
    "block_hash" varchar NOT NULL,
    "log_index" varchar NOT NULL,
    "removed" boolean NOT NULL
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("block_hash") REFERENCES "blocks" ("block_hash");
ALTER TABLE "transactions" ADD FOREIGN KEY ("block_number") REFERENCES "blocks" ("block_number");

ALTER TABLE "logs" ADD FOREIGN KEY ("block_number") REFERENCES "blocks" ("block_number");

ALTER TABLE "logs" ADD FOREIGN KEY ("transaction_hash") REFERENCES "transactions" ("transaction_hash");
ALTER TABLE "logs" ADD FOREIGN KEY ("transaction_index") REFERENCES "transactions" ("transaction_index");

CREATE INDEX ON "blocks" ("block_number");
CREATE INDEX ON "blocks" ("block_hash");

CREATE INDEX ON "transactions" ("block_hash");
CREATE INDEX ON "transactions" ("block_number");
CREATE INDEX ON "transactions" ("transaction_hash");
