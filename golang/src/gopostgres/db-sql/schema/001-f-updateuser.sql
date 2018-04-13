-- example usage
-- select updateUser(id := 1, username := 'adhuliya');

CREATE OR REPLACE FUNCTION updateUser (
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
    username        = COALESCE(updateUser.username, tt.username),
    password        = COALESCE(updateUser.password, tt.password),
    nickname        = COALESCE(updateUser.nickname, tt.nickname),
    fullname        = COALESCE(updateUser.fullname, tt.fullname),
    mobile          = COALESCE(updateUser.mobile,   tt.mobile),
    email           = COALESCE(updateUser.email,    tt.email),
    role            = COALESCE(updateUser.role,     tt.role),
    status          = COALESCE(updateUser.status,   tt.status),
    modifiedOn      = NOW()
    WHERE tt.id = updateUser.id;
    RETURN FOUND;
END;
$$;


