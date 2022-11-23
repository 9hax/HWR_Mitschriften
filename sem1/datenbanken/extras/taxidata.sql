DROP TABLE IF EXISTS Firma;
DROP TABLE IF EXISTS Auto;

CREATE TABLE Firma (
    FName VARCHAR2(10),
    /* Eigentlich TEXT */
    AnzahlMA INTEGER,
    IBAN TEXT
    /* SQLite3 akzeptiert viele ähnliche Datentypen */
);
CREATE TABLE Auto (
    KFZZeichen VARCHAR2(10),
    Modell VARCHAR2(8),
    Baujahr INTEGER
);

DELETE FROM Firma;
DELETE FROM Auto;
INSERT INTO Firma
VALUES ('Maxi-Taxi', 350, 'DE1234'),
    ('Luxi-Taxi', 90, 'DE2345'),
    ('Fixi-Taxi', 120, 'FR1234');
INSERT INTO Auto
VALUES ('BAZ-1718', 'VW', 2011),
    ('BMP-1718', 'Skoda', 2015),
    ('BKA-4253', 'Mercedes', 2000),
    ('BAZ-9876', 'BMW', 2011),
    ('BAZ-6789', 'VW', 2011);
DROP TABLE IF EXISTS betreibtAuto;
CREATE TABLE betreibtAuto (
    FName VARCHAR2(10),
    KFZZeichen VARCHAR2(10),
    FOREIGN KEY(FName) REFERENCES Firma(FName),
    FOREIGN KEY(KFZZeichen) REFERENCES Auto(KFZZeichen)
);
DELETE FROM betreibtAuto;
INSERT INTO betreibtAuto
VALUES ('Maxi-Taxi', 'BAZ-1718'),
    ('Maxi-Taxi', 'BMP-1718'),
    ('Luxi-Taxi', 'BKA-4253'),
    ('Luxi-Taxi', 'BAZ-9876'),
    ('Fixi-Taxi', 'BAZ-6789');

SELECT *
FROM Auto a, Firma f, betreibtAuto b
WHERE f.FName = b.FName AND b.KFZZeichen = a.KFZZeichen