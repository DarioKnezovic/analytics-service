-- Migration created at 2023-11-18T19:45:16Z

ALTER TABLE modal_ctr_tracking ADD COLUMN new_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE modal_ctr_tracking DROP COLUMN timestamp;
ALTER TABLE modal_ctr_tracking RENAME COLUMN new_timestamp TO timestamp;
