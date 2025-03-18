import mysql.connector
import yaml

# Load database config
with open("config.yaml", "r") as file:
    config = yaml.safe_load(file)

# Connect to MySQL
db = mysql.connector.connect(
    host=config["db"]["host"],
    user=config["db"]["user"],
    password=config["db"]["password"],
    database=config["db"]["database"]
)
cursor = db.cursor()

# Function to add an item to the shopping list
def add_to_shopping_list(user_id, ingredient_id, quantity):
    query = """
    INSERT INTO shopping_list_items (user_id, ingredient_id, quantity)
    VALUES (%s, %s, %s)
    ON DUPLICATE KEY UPDATE quantity = quantity + VALUES(quantity)
    """
    cursor.execute(query, (user_id, ingredient_id, quantity))
    db.commit()
    print("Item added to shopping list.")
