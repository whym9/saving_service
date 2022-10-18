CREATE DATABASE db;

USE db;

CREATE TABLE Pcap_Files (
    ID INT UNIQUE,
    Path VARCHAR(256),
    INDEX (ID)
);

CREATE TABLE Protocols (
    ID INT UNIQUE,
    Name UNIQUE VARCHAR(256), 
    INDEX (ID) 
);


CREATE TABLE File_Statistic (
    FileID INT,
    ProtocolID INT, 
    FOREIGN KEY (FileID) REFERENCES Pcap_Files(ID),
    FOREIGN KEY (ProtocolID) REFERENCES Protocols(ID),
    INDEX(FileID, ProtocolID),
    Count INT
);

INSERT Protocols(Name) VALUE('TCP');
INSERT Protocols(Name) VALUE('UDP');
INSERT Protocols(Name) VALUE('IPv4');
INSERT Protocols(Name) VALUE('IPv6');
