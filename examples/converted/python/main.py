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
    return int(time.time_ns() % 10000)


def save_user(user):
    print(f"Saving user: {user.name}")
    return None


def get_user_by_id(id):
    if id < 0:
        raise ValueError(f"Invalid user ID: {id}")

    return User(id=id, name="Example User", email="user@example.com", created_at=time.time())


def main():
    user = new_user("John Doe", "john@example.com")

    if save_user(user) is not None:
        print(f"Error saving user")
        sys.exit(1)

    try:
        retrieved_user = get_user_by_id(user.id)
    except ValueError as err:
        print(f"Error retrieving user: {err}")
        sys.exit(1)

    print("User Information:")
    print(f"ID: {retrieved_user.id}")
    print(f"Name: {retrieved_user.name}")
    print(f"Email: {retrieved_user.email}")
    print(f"Created: {time.strftime('%Y-%m-%dT%H:%M:%S', time.gmtime(retrieved_user.created_at))}")


if __name__ == "__main__":
    main()
