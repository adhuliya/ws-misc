-- example usage
-- select * from user_Select();
-- never return NULL values (hard to handle them in golang)

CREATE OR REPLACE FUNCTION user_Select (
    ioffset     INTEGER DEFAULT 0,
    ilimit      INTEGER DEFAULT 100
) RETURNS TABLE (
    id          INTEGER,
    username    TEXT,
    password    TEXT,
    nickname    TEXT,
    fullname    TEXT,
    mobile      TEXT,
    email       TEXT,
    role        TEXT,
    status      TEXT,
    createdOn   TIMESTAMPTZ,
    modifiedOn  TIMESTAMPTZ
)
LANGUAGE plpgsql SECURITY DEFINER AS $$

BEGIN
    RETURN QUERY SELECT
    tt.id,
    COALESCE(tt.username, ''), -- avoiding NULL
    COALESCE(tt.password, ''), -- avoiding NULL
    COALESCE(tt.nickname, ''), -- avoiding NULL
    COALESCE(tt.fullname, ''), -- avoiding NULL
    COALESCE(tt.mobile, ''),   -- avoiding NULL
    COALESCE(tt.email, ''),    -- avoiding NULL
    COALESCE(tt.role, ''),     -- avoiding NULL
    COALESCE(tt.status, ''),   -- avoiding NULL
    COALESCE(tt.createdOn, to_timestamp(0)), -- avoiding NULL
    COALESCE(tt.modifiedOn, to_timestamp(0)) -- avoiding NULL
    FROM
    public.user tt
    ORDER BY tt.id
    LIMIT selectUser.ilimit
    OFFSET selectUser.ioffset;
END;
$$;


