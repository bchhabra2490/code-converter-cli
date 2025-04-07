class User {
  constructor(id, name, email, createdAt) {
    this.id = id;
    this.name = name;
    this.email = email;
    this.createdAt = createdAt;
  }
}

// Generates a simple random ID
function generateID() {
  return Math.floor(Date.now() % 10000);
}

// Creates a new user
function newUser(name, email) {
  return new User(generateID(), name, email, new Date());
}

// Simulates saving a user to a database
function saveUser(user) {
  console.log(`Saving user: ${user.name}`);
  return Promise.resolve();
}

// Retrieves a user by ID
function getUserByID(id) {
  if (id < 0) {
    return Promise.reject(new Error(`invalid user ID: ${id}`));
  }

  // Mock data
  return Promise.resolve(
    new User(id, "Example User", "user@example.com", new Date())
  );
}

async function main() {
  try {
    // Create a new user
    const user = newUser("John Doe", "john@example.com");

    // Save the user
    await saveUser(user);

    // Retrieve the user
    const retrievedUser = await getUserByID(user.id);

    // Display user information
    console.log("User Information:");
    console.log(`ID: ${retrievedUser.id}`);
    console.log(`Name: ${retrievedUser.name}`);
    console.log(`Email: ${retrievedUser.email}`);
    console.log(`Created: ${retrievedUser.createdAt.toISOString()}`);
  } catch (error) {
    console.error(`Error: ${error.message}`);
    process.exit(1);
  }
}

main();