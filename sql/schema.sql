CREATE TABLE contents (
    uuid VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    contentType VARCHAR(100),
    categories BLOB,
    tags BLOB,
    author VARCHAR(100),
    publicationDate DATETIME,
    contentUrl VARCHAR(255),
    duration INT,
    language VARCHAR(50),
    coverImage VARCHAR(255),
    metadata BLOB,
    status VARCHAR(50),
    source VARCHAR(100),
    visibility VARCHAR(50)
);
