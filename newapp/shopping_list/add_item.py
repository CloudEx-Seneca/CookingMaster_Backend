from flask import Flask, request, jsonify
import mysql.connector
import yaml

# Load database config
with open("shop_config.yaml", "r") as file:
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
@app.route("/add_item", methods=["POST"])
def add_to_shopping_list(user_id, ingredient_id):
    query = """
    INSERT INTO shopping_list_items (user_id, ingredient_id)
    VALUES (%s, %s)
    """
    cursor.execute(query, (user_id, ingredient_id))
    db.commit()
    print("Item added to shopping list.")

if __name__ == "__main__":
    app.run(debug=True)
