CREATE DATABASE db;

USE db;

CREATE TABLE Pcap_Files (
    ID INT INDEX,
    Path VARCHAR(256),
    PRIMARY KEY (ID)
);

CREATE TABLE Protocols (
    ID INT INDEX,
    Name UNIQUE VARCHAR(256), 
    PRIMARY KEY (ID) 
);


CREATE TABLE File_Statistic (
    FileID INT,
    ProtocolID INT, 
    FOREIGN KEY (FileID) REFERENCES Pcap_Files(ID),
    FOREIGN KEY (ProtocolID) REFERENCES Protocols(ID),
    Count INT
);
