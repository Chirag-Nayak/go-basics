DROP TABLE if exists public.account_info;
CREATE TABLE account_info (
    id bigserial PRIMARY KEY,
    account_name character varying(100) NOT NULL CONSTRAINT account_name_must_be_unique UNIQUE,
    currency_name character varying(10) NOT NULL,
    balance decimal(16,8)
);

comment on column account_info.balance is 'Can store balance upto 99999999.99999999';

INSERT INTO account_info (account_name, currency_name, balance) VALUES
  ('My Account 1', 'USD',10000.12345678),
  ('My Account 2', 'USD',99999999.99999999),
  ('My Account 3', 'USD',-99999999.99999999);
