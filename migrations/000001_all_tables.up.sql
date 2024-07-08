CREATE TABLE users
(
    id                  serial       not null unique,
    name                VARCHAR(255) NOT NULL,
    username            VARCHAR(255) NOT NULL unique,
    password            VARCHAR(255) NOT NULL,
    bookmarked_pets_ids INTEGER[]
);

CREATE TABLE pets
(
    id          UUID PRIMARY KEY,
    name        VARCHAR(31) NOT NULL,
    type        VARCHAR(31) NOT NULL,
    breed       VARCHAR(31),
    age         INTEGER     NOT NULL,
    size        VARCHAR(31),
    gender      VARCHAR(31) NOT NULL,
    description VARCHAR(255),
    photo_url   VARCHAR(255),
    add_date    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);