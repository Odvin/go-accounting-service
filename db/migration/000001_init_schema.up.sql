-- Client schema
CREATE SCHEMA IF NOT EXISTS client;
-- Profile
CREATE TYPE ADMINISTRATIVE_STATUS AS ENUM (
  'adm:active', 'adm:blocked', 'adm:suspended', 'adm:processed'
);
CREATE TYPE KYC_STATUS AS ENUM (
  'kyc:unconfirmed',
  'kyc:confirmed',
  'kyc:pending',
  'kyc:rejected',
  'kyc:resubmission',
  'kyc:initiated'
);
CREATE TABLE IF NOT EXISTS client.profile (
  id uuid PRIMARY KEY,
  adm ADMINISTRATIVE_STATUS NOT NULL,
  kyc KYC_STATUS NOT NULL,
  name VARCHAR(15) NOT NULL,
  surname VARCHAR(30) NOT NULL,
  updated TIMESTAMPTZ DEFAULT Now() NOT NULL,
  created TIMESTAMPTZ DEFAULT Now() NOT NULL
);
COMMENT ON TABLE client.profile IS 'Account holders profile';
COMMENT ON COLUMN client.profile.adm IS 'Client status set by the system admin';
COMMENT ON COLUMN client.profile.kyc IS 'Client status from security service (Know Your Client)';
