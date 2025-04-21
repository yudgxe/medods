-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth (
    uuid           CHAR(128) PRIMARY KEY,
    refresh_token  TEXT CHECK(refresh_token <> '')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth;
-- +goose StatementEnd
