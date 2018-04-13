CREATE TABLE IF NOT EXISTS public.user (
    id          SERIAL  NOT NULL,
    username    TEXT UNIQUE NOT NULL,
    password    TEXT,
    nickname    TEXT,
    fullname    TEXT,
    mobile      TEXT UNIQUE NOT NULL,
    email       TEXT UNIQUE NOT NULL,
    role        TEXT NOT NULL,
    status      TEXT NOT NULL,
    createdOn   TIMESTAMPTZ NOT NULL,
    modifiedOn  TIMESTAMPTZ,
    CONSTRAINT userinfo_pkey PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION update_user (
    id          serial,
    username    text,
    password    text,
    nickname    text,
    fullname    text,
    mobile      text,
    email       text,
    role        text,
    status      text,
    createdOn   TIMESTAMPTZ,
    modifiedOn  TIMESTAMPTZ
) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$
BEGIN
    UPDATE public.user tt
    SET 
    username        = COALESCE(update_user.username, tt.username),
    password        = COALESCE(update_user.password, tt.password),
    nickname        = COALESCE(update_user.nickname, tt.nickname),
    fullname        = COALESCE(update_user.fullname, tt.fullname),
    mobile          = COALESCE(update_user.mobile,   tt.mobile),
    email           = COALESCE(update_user.email,    tt.email),
    role            = COALESCE(update_user.role,     tt.role),
    status          = COALESCE(update_user.status,   tt.status),
    modifiedOn      = NOW()
    WHERE tt.id = update_user.id
    RETURN FOUND;
END;
$$;


