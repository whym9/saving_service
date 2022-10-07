CREATE DATABASE pcap_files;

USE pcap_files;

CREATE TABLE File_Statistics (
    FilePath VARCHAR(256),
    ProtocolTCP INT,
    UDP INT,
    IPv4 INT,
    IPv6 INT,
    PRIMARY KEY (FilePath)
    );
