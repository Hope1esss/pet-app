CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    google_user_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
bookmarked_pets_ids INTEGER[]
);

CREATE TABLE pets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(31) NOT NULL,
    type VARCHAR(31) NOT NULL,
    breed VARCHAR(31),
    age INTEGER NOT NULL,
    size VARCHAR(31),
    gender VARCHAR(31),
    description VARCHAR(255),
    photo_url VARCHAR(255)
);