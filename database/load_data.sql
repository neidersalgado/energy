LOAD DATA LOCAL INFILE '/docker-entrypoint-initdb.d/test_bia.csv'
INTO TABLE energy
FIELDS TERMINATED BY ','
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;