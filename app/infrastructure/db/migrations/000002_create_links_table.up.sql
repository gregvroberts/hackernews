CREATE TABLE IF NOT EXISTS Links(
    ID SERIAL PRIMARY KEY,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID SERIAL REFERENCES Users(ID)
);