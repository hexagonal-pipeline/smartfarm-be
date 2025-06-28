-- migrate:up
ALTER TABLE farm_plots ADD COLUMN persona_prompt TEXT;

-- migrate:down
ALTER TABLE farm_plots DROP COLUMN persona_prompt;

