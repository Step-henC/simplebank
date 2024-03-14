ALTER TABLE IF EXISTS "account" DROP CONSTRAINT IF EXISTS "owner_currency_key";

ALTER TABLE IF EXISTS "account" DROP CONSTRAINT IF EXISTS "account_owner_fkey";
-- found account owner fkey in info section of table plus
DROP TABLE IF EXISTS "user";
