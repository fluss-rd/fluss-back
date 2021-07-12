// Create database
db = new Mongo().getDB("flussDB");

// Collections
db.createCollection("roles");
db.createCollection("users");
db.createCollection("modules");
db.createCollection("rivers");

// Unique indexes fo
db.roles.createIndex({ roleName: 1 }, { unique: true });
db.users.createIndex({ email: 1 }, { unique: true });
db.modules.createIndex({ riverID: 1,  alias: 1}, { unique: true });
db.modules.createIndex({ phoneNumber: 1}, { unique: true });
db.rivers.createIndex({ name: 1 }, { unique: true });

db.roles.insert({
    roleName: "superAdmin",
    permissions: [
        {
            resource: "*",
            actions: ["*"],
        }
    ]
});

// This is not secure, figure out a way to do it
db.users.insert({
    _id: "unique-id",
    email: "fluss.rd.admin@gmai.com",
    password: "$2b$10$W8aSBwYLOANnCYUNEhM.runbR4fs7jBT5OPOWJkt7ShddmBfMUxvS",
    roleName: "superAdmin" 
});

