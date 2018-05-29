CREATE TABLE IF NOT EXISTS public.user (
    id          SERIAL  NOT NULL,
    username    TEXT UNIQUE NOT NULL,
    password    TEXT DEFAULT 'p@ssw0rd',
    nickname    TEXT DEFAULT 'HOPPER',
    fullname    TEXT DEFAULT 'HOPPER',
    mobile      TEXT UNIQUE NOT NULL,
    email       TEXT UNIQUE NOT NULL,
    role        TEXT NOT NULL DEFAULT 'user1',
    status      TEXT NOT NULL DEFAULT 'disabled',
    createdOn   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modifiedOn  TIMESTAMPTZ,
    CONSTRAINT userinfo_pkey PRIMARY KEY (id)
);

