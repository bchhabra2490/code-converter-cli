import time
import sys


class User:
    def __init__(self, id, name, email, created_at):
        self.id = id
        self.name = name
        self.email = email
        self.created_at = created_at


def new_user(name, email):
    return User(id=generate_id(), name=name, email=email, created_at=time.time())


def generate_id():
    return int(time.time() * 1000) % 10000


def save_user(user):
    # Simulate database operation
    print(f"Saving user: {user.name}")


def get_user_by_id(id):
    # Simulate database lookup
    if id < 0:
        raise ValueError(f"invalid user ID: {id}")

    # Mock data
    return User(id=id, name="Example User", email="user@example.com", created_at=time.time())


def main():
    # Create a new user
    user = new_user("John Doe", "john@example.com")

    # Save the user
    try:
        save_user(user)
    except Exception as err:
        print(f"Error saving user: {err}", file=sys.stderr)
        sys.exit(1)

    # Retrieve the user
    try:
        retrieved_user = get_user_by_id(user.id)
    except Exception as err:
        print(f"Error retrieving user: {err}", file=sys.stderr)
        sys.exit(1)

    # Display user information
    print("User Information:")
    print(f"ID: {retrieved_user.id}")
    print(f"Name: {retrieved_user.name}")
    print(f"Email: {retrieved_user.email}")
    print("Created:", time.strftime("%Y-%m-%dT%H:%M:%S", time.localtime(retrieved_user.created_at)))


if __name__ == "__main__":
    main()
