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
    id          serial       not null unique PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    type        VARCHAR(255) NOT NULL,
    breed       VARCHAR(255),
    age         VARCHAR(255) NOT NULL,
    size        VARCHAR(255),
    gender      VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    add_date    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    added_by_user_id    INTEGER NOT NULL
);