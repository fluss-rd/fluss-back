// Create database
db = new Mongo().getDB("flussDB");

// Collections
db.createCollection("roles");
db.createCollection("users");

// Unique indexes
db.roles.createIndex({ roleName: 1 }, { unique: true });
db.users.createIndex({ email: 1 }, { unique: true });