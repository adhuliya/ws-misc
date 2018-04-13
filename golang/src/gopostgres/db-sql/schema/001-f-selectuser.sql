-- example usage
-- select * from selectUser();
-- never return NULL values (hard to handle them in golang)

CREATE OR REPLACE FUNCTION selectUser (
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
    COALESCE(tt.username, ''),
    COALESCE(tt.password, ''),
    COALESCE(tt.nickname, ''), -- avoiding NULL
    COALESCE(tt.fullname, ''),
    COALESCE(tt.mobile, ''),
    COALESCE(tt.email, ''),
    COALESCE(tt.role, ''),
    COALESCE(tt.status, ''),
    COALESCE(tt.createdOn, to_timestamp(0)),
    COALESCE(tt.modifiedOn, to_timestamp(0))
    FROM
    public.user tt
    ORDER BY tt.id
    LIMIT selectUser.ilimit
    OFFSET selectUser.ioffset;
END;
$$;


