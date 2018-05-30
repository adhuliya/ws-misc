-- example usage

-- select user_Insert(username := 'rimi', mobile := '+918756268826', email := 'rimi@gmail.com');

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
        insertuser.username,
        COALESCE(insertUser.password, 'p@ssw0rd'),
        COALESCE(insertUser.nickname, 'HOPPER'),
        COALESCE(insertUser.fullname, 'HOPPER'),
        insertUser.mobile,
        insertUser.email,
        COALESCE(insertUser.role, 'user1'),
        COALESCE(insertUser.status, 'disabled'),
        COALESCE(insertUser.createdOn, CURRENT_TIMESTAMP),
        NULL
    )
    RETURNING id INTO insertUser.insertId;
    --RETURN FOUND;
END;
$$;


