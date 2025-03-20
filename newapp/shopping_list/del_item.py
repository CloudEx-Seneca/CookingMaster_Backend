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

# Function to delete an ingredient from the shopping list
def del_item(user_id, ingredient_id):
    query = "DELETE FROM shopping_list_items WHERE user_id = %s AND ingredient_id = %s"
    cursor.execute(query, (user_id, ingredient_id))
    db.commit()
    print("Ingredient removed from shopping list.")
