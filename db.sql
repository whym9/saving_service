CREATE DATABASE db;

USE db;

CREATE TABLE Pcap_Files (
    FileID INT UNIQUE,
    FilePath VARCHAR(256) UNIQUE,
    PRIMARY KEY (FileID)
);

CREATE TABLE Protocols (
    ProtocolName VARCHAR(256) UNIQUE, 
    PRIMARY KEY (ProtocolName) 
);


CREATE TABLE File_Statistic (
    FilePath VARCHAR (256) UNIQUE,
    ProtocolName VARCHAR(256) UNIQUE, 
    FOREIGN KEY (FilePath) REFERENCES Pcap_Files(FilePath) ON DELETE CASCADE,
    FOREIGN KEY (ProtocolName) REFERENCES Protocols(ProtocolName),
    PRIMARY KEY (FilePath, ProtocolName),
    Count INT
);