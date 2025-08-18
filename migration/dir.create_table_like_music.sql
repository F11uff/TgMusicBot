CREATE TABLE IF NOT EXISTS LikeMusic (
--     Id INT PRIMARY KEY,
    IdMusic SERIAL PRIMARY KEY ,
    Music VARCHAR NOT NULL,
    Artist VARCHAR NOT NULL
);

INSERT INTO LikeMusic(IdMusic, Artist, Music) VALUES (1, 'Ed Sheeren', 'Perfect'),
                                                (2, 'Lady Gaga', 'Abracadabra')