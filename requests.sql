SELECT Name FROM Protocols; # A request to get all protocol names.

SELECT # Get a table with all protol names and their counts given file path.
	P.Name, 
    S.Count 
FROM Pcap_Files F 
INNER JOIN File_Statistic S 
ON F.ID = S.FileID 
INNER JOIN Protocols P 
ON S.ProtocolID = P.ID 
WHERE F.Path='file1';

SELECT  # Get a table with all protol names and their counts given file ID.
	P.Name, 
    S.Count 
FROM Pcap_Files F 
INNER JOIN File_Statistic S 
ON F.ID = S.FileID 
INNER JOIN Protocols P 
ON S.ProtocolID = P.ID 
WHERE F.ID = 2;

SELECT F.Path # Get all the files that have given protocol.
FROM Pcap_Files F 
INNER JOIN File_Statistic S
ON F.ID = S.FileID
INNER JOIN Protocols P 
ON S.ProtocolID = P.ID 
WHERE P.Name = 'udp';

SELECT # get the sum of all the packets with the given protocol
	P.Name,
	SUM(S.Count)
FROM Protocols P 
INNER JOIN File_Statistic S
ON P.ID = S.ProtocolID
WHERE P.Name = 'tcp';

SELECT # get the amount of all the packets in a file given its name.
	F.Path,
	SUM(S.Count)
FROM Pcap_Files F 
INNER JOIN File_Statistic S
ON F.ID = S.FileID
INNER JOIN Protocols P
ON S.ProtocolID = P.ID
WHERE F.Path = 'file5' AND (P.Name = 'tcp' OR P.Name = 'udp');
