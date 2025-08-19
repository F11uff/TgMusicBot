CREATE TABLE IF NOT EXISTS LikeMusic (
--     Id INT PRIMARY KEY,
    IdMusic SERIAL PRIMARY KEY ,
    Music VARCHAR,
    Artist VARCHAR
);

INSERT INTO LikeMusic(Artist, Music) VALUES ('Ed Sheeren', 'Perfect'),
                                                ('Lady Gaga', 'Abracadabra')