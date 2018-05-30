-- example usage
-- select user_Delete(id := 1);

CREATE OR REPLACE FUNCTION user_Delete (
    id          INTEGER
) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$

BEGIN
    DELETE FROM public.user tt
    WHERE 
    tt.id = deleteUser.id;
    RETURN FOUND;
END;
$$;


