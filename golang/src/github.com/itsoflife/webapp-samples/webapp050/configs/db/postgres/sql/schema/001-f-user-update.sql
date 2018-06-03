-- example usage
-- select user_Update(id := 1, username := 'adhuliya');
\c hop;

CREATE OR REPLACE FUNCTION user_Update (
    id          INTEGER,
    username    TEXT DEFAULT NULL,
    password    TEXT DEFAULT NULL,
    nickname    TEXT DEFAULT NULL,
    fullname    TEXT DEFAULT NULL,
    mobile      TEXT DEFAULT NULL,
    email       TEXT DEFAULT NULL,
    role        TEXT DEFAULT NULL,
    status      TEXT DEFAULT NULL,
    createdOn   TIMESTAMPTZ DEFAULT NULL,
    modifiedOn  TIMESTAMPTZ DEFAULT NULL
) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$

BEGIN
    UPDATE public.user tt
    SET 
    username        = COALESCE(user_Update.username, tt.username),
    password        = COALESCE(user_Update.password, tt.password),
    nickname        = COALESCE(user_Update.nickname, tt.nickname),
    fullname        = COALESCE(user_Update.fullname, tt.fullname),
    mobile          = COALESCE(user_Update.mobile,   tt.mobile),
    email           = COALESCE(user_Update.email,    tt.email),
    role            = COALESCE(user_Update.role,     tt.role),
    status          = COALESCE(user_Update.status,   tt.status),
    modifiedOn      = NOW()
    WHERE tt.id = user_Update.id;
    RETURN FOUND;
END;
$$;


