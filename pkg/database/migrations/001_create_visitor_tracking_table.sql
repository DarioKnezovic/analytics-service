-- Migration created at 2023-11-04T23:41:09Z

CREATE TABLE IF NOT EXISTS visitor_tracking (
                                visitor_id INT IDENTITY(1,1) PRIMARY KEY,
                                timestamp TIMESTAMP NOT NULL,
                                adblock_user BOOLEAN NOT NULL,
                                campaign_id INT
);