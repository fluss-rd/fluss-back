// Collections
db.createCollection("roles");
db.createCollection("users");

// Unique indexes
db.roles.createIndex({ roleName: 1 }, { unique: true });