// Create database
db = new Mongo().getDB("flussDB");

// Collections
db.createCollection("roles");
db.createCollection("users");

// Unique indexes
db.roles.createIndex({ roleName: 1 }, { unique: true });
db.users.createIndex({ email: 1 }, { unique: true });

db.roles.insert({
    roleName: "superAdmin",
    permissions: [
        {
            resource: "*",
            action: "*",
        }
    ]
});

// This is not secure, figure out a way to do it
db.users.insert({
    _id: "unique-id",
    email: "admin@admin",
    password: "$2b$10$W8aSBwYLOANnCYUNEhM.runbR4fs7jBT5OPOWJkt7ShddmBfMUxvS",
    roleName: "superAdmin" 
});

