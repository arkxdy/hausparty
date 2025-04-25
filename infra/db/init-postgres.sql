-- 1) Create service databases
CREATE DATABASE identity_db;   -- holds users, auth_credentials, sessions, roles, user_roles
CREATE DATABASE party_db;      -- holds parties, rsvps, etc.
CREATE DATABASE admin_db;      -- holds admin_actions, logs, etc.

-- 2) Add UUID extension to each (needed for uuid_generate_v4)
\connect identity_db
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\connect party_db
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\connect admin_db
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
