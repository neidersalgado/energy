CREATE TABLE IF NOT EXISTS energy (
    id VARCHAR(36),
    meter_id INT,
    active_energy DOUBLE,
    reactive_energy DOUBLE,
    capacitive_reactive DOUBLE,
    solar DOUBLE,
    date DATETIME
    );