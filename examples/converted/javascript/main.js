class User {
  constructor(id, name, email, createdAt) {
    this.id = id;
    this.name = name;
    this.email = email;
    this.createdAt = createdAt;
  }
}

function NewUser(name, email) {
  return new User(generateID(), name, email, new Date());
}

function generateID() {
  return Math.floor(Date.now() % 10000);
}

function SaveUser(user) {
  return new Promise((resolve, reject) => {
    // Simulate database operation
    console.log(`Saving user: ${user.name}`);
    resolve();
  });
}

function GetUserByID(id) {
  return new Promise((resolve, reject) => {
    if (id < 0) {
      reject(new Error(`invalid user ID: ${id}`));
    } else {
      // Mock data
      resolve(new User(id, "Example User", "user@example.com", new Date()));
    }
  });
}

async function main() {
  try {
    // Create a new user
    const user = NewUser("John Doe", "john@example.com");

    // Save the user
    await SaveUser(user);

    // Retrieve the user
    const retrievedUser = await GetUserByID(user.id);

    // Display user information
    console.log("User Information:");
    console.log(`ID: ${retrievedUser.id}`);
    console.log(`Name: ${retrievedUser.name}`);
    console.log(`Email: ${retrievedUser.email}`);
    console.log(`Created: ${retrievedUser.createdAt.toISOString()}`);
  } catch (err) {
    console.error(`Error: ${err.message}`);
    process.exit(1);
  }
}

main();