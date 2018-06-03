-- example usage

-- select user_Insert(username := 'rimi', mobile := '+918756268826', email := 'rimi@gmail.com');
\c hop;

CREATE OR REPLACE FUNCTION user_Insert (
    username    TEXT,
    mobile      TEXT,
    email       TEXT,
    password    TEXT DEFAULT NULL,
    nickname    TEXT DEFAULT NULL,
    fullname    TEXT DEFAULT NULL,
    role        TEXT DEFAULT NULL,
    status      TEXT DEFAULT NULL,
    createdOn   TIMESTAMPTZ DEFAULT NULL,
    modifiedOn  TIMESTAMPTZ DEFAULT NULL,
    OUT insertId INTEGER
) LANGUAGE plpgsql SECURITY DEFINER

AS $$
BEGIN
    insert into public.user (
        -- id,
        username,
        password,
        nickname,
        fullname,
        mobile,
        email,
        role,
        status,
        createdOn,
        modifiedOn
    )

    VALUES
    (
        user_Insert.username,
        COALESCE(user_Insert.password, 'p@ssw0rd'),
        COALESCE(user_Insert.nickname, 'HOPPER'),
        COALESCE(user_Insert.fullname, 'HOPPER'),
        user_Insert.mobile,
        user_Insert.email,
        COALESCE(user_Insert.role, 'user1'),
        COALESCE(user_Insert.status, 'disabled'),
        COALESCE(user_Insert.createdOn, CURRENT_TIMESTAMP),
        NULL
    )
    RETURNING id INTO user_Insert.insertId;
    --RETURN FOUND;
END;
$$;


