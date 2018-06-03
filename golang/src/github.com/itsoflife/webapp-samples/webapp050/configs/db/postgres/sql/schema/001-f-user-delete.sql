-- example usage
-- select user_Delete(id := 1);
\c hop;

CREATE OR REPLACE FUNCTION user_Delete (
    id          INTEGER
) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$

BEGIN
    DELETE FROM public.user tt
    WHERE 
    tt.id = user_Delete.id;
    RETURN FOUND;
END;
$$;


