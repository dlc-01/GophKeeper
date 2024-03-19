-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

