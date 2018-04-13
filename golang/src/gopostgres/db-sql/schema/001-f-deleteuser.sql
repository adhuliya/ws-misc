-- example usage
-- select updateUser(id := 1, username := 'adhuliya');

CREATE OR REPLACE FUNCTION deleteUser (
    id          INTEGER
) RETURNS BOOLEAN LANGUAGE plpgsql SECURITY DEFINER AS $$

BEGIN
    DELETE FROM public.user tt
    WHERE 
    tt.id = deleteUser.id;
    RETURN FOUND;
END;
$$;


