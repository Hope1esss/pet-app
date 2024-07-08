CREATE TABLE users
(
    id                  serial       not null unique PRIMARY KEY,
    name                VARCHAR(255) NOT NULL,
    username            VARCHAR(255) NOT NULL unique,
    password            VARCHAR(255) NOT NULL,
    bookmarked_pets_ids INTEGER[]
);

CREATE TABLE pets
(
    id          serial      not null unique PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    type        VARCHAR(255) NOT NULL,
    breed       VARCHAR(255),
    age         VARCHAR(255)     NOT NULL,
    size        VARCHAR(31),
    gender      VARCHAR(31) NOT NULL,
    description VARCHAR(255),
    add_date    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);