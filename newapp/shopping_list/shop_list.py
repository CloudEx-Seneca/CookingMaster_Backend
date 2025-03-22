from flask import Flask, request, jsonify
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

app = Flask(__name__)

@app.route("/add_item", methods=["POST"])
def add_to_shopping_list():
    data = request.json
    user_id = data["user_id"]
    ingredient = data["ingredient"]

    # Use the existing db connection to get the cursor
    cursor = db.cursor()

    query = """
    INSERT INTO shopping_list_items (user_id, ingredient)
    VALUES (%s, %s)
    ON DUPLICATE KEY UPDATE ingredient = VALUES(ingredient)
    """
    cursor.execute(query, (user_id, ingredient))
    db.commit()
    
    # Close the cursor after executing the query
    cursor.close()

    return jsonify({"message": "Item added to shopping list"}), 201


# Get shopping list for a user
@app.route("/get_list/<int:user_id>", methods=["GET"])
def get_shopping_list(user_id):
    query = """
    SELECT ingredient FROM shopping_list_items WHERE user_id = %s
    """

    cursor.execute(query, (user_id,))
    items = [item[0] for item in cursor.fetchall()]
    return jsonify(items)


# Remove an item from the shopping list
@app.route("/remove_item", methods=["POST"])
def remove_item():
    data = request.json
    user_id = data["user_id"]
    ingredient = data["ingredient"]

    query = """
    DELETE FROM shopping_list_items
    WHERE user_id = %s AND ingredient = %s
    """
    cursor.execute(query, (user_id, ingredient))
    db.commit()
    return jsonify({"message": "Item removed from shopping list"}), 200

if __name__ == "__main__":
    app.run(debug=True)
