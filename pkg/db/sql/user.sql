---------------------------------------------------------------------
-- query: user_insert
---------------------------------------------------------------------
INSERT INTO users (
        username,
        email,
        email_verified,
        password_hash,
        first_name,
        last_name,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING user_id;
---------------------------------------------------------------------
-- query: user_select_by_email
---------------------------------------------------------------------
SELECT user_id,
    username,
    email,
    email_verified,
    password_hash,
    first_name,
    last_name,
    created_at,
    updated_at
FROM users
WHERE email = $1;
---------------------------------------------------------------------
