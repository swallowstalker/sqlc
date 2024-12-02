CREATE TABLE IF NOT EXISTS auth_user
(
    id      SERIAL  PRIMARY KEY, -- ID
    name    VARCHAR NOT NULL,
    user_type VARCHAR NOT NULL,
    extent  JSONB   NOT NULL DEFAULT '{}', --
    dept_id INTEGER              -- 部门ID
);