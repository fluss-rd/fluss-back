// Collections
db.createCollection("roles");

// Unique indexes
db.roles.createIndex({ roleName: 1 }, { unique: true });