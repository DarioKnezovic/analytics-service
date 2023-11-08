-- Migration created at 2023-11-08T19:07:21Z

CREATE TABLE modal_ctr_tracking (
                                    id INT IDENTITY(1,1) PRIMARY KEY,
                                    session VARCHAR(255) NOT NULL,
                                    timestamp TIMESTAMP NOT NULL,
                                    interaction_type VARCHAR(50) NOT NULL,
                                    object_id VARCHAR(255) NOT NULL,
                                    additional_data VARCHAR(255),
                                    campaign_id INT NOT NULL
);