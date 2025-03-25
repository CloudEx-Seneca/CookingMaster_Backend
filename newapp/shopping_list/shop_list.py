from flask import Flask, request, jsonify
import mysql.connector
from mysql.connector import errorcode
import yaml

# Load database config
with open("config.yaml", "r") as file:
    config = yaml.safe_load(file)

# Connect to MySQL
try:
    db = mysql.connector.connect(
        host=config["db"]["host"],
        user=config["db"]["user"],
        password=config["db"]["password"],
        database=config["db"]["database"]
    )
except mysql.connector.Error as err:
    if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
        print("Something is wrong with your user name or password")
    elif err.errno == errorcode.ER_BAD_DB_ERROR:
        print("Database does not exist")
    else:
        print(err)
    exit()

app = Flask(__name__)

# Create a cursor before every request
@app.before_request
def create_cursor():
    global cursor
    cursor = db.cursor()

# Close the cursor after every request
@app.teardown_request
def close_cursor(exception=None):
    global cursor
    if cursor:
        cursor.close()

@app.route("/add_item", methods=["POST"])
def add_to_shopping_list():
    data = request.json
    user_id = data.get("user_id")
    ingredients = data.get("ingredients")

    if not user_id or not ingredients or not isinstance(ingredients, list):
        return jsonify({"message": "User ID and a list of ingredients are required"}), 400

    try:
        for ingredient in ingredients:
            query = """
            INSERT INTO shopping_list_items (user_id, ingredient)
            VALUES (%s, %s)
            ON DUPLICATE KEY UPDATE ingredient = VALUES(ingredient)
            """
            cursor.execute(query, (user_id, ingredient))
        
        db.commit()
        return jsonify({"message": "Items added to shopping list"}), 201
    except mysql.connector.Error as err:
        db.rollback()
        return jsonify({"message": f"Error: {err}"}), 500

@app.route("/get_list/<int:user_id>", methods=["GET"])
def get_shopping_list(user_id):
    try:
        query = """
        SELECT ingredient FROM shopping_list_items WHERE user_id = %s
        """
        cursor.execute(query, (user_id,))
        items = [item[0] for item in cursor.fetchall()]
        
        if not items:
            return jsonify({"message": "No items found for this user"}), 404
        
        return jsonify(items), 200
    except mysql.connector.Error as err:
        return jsonify({"message": f"Error: {err}"}), 500

@app.route("/remove_item", methods=["POST"])
def remove_item():
    data = request.json
    user_id = data.get("user_id")
    ingredients = data.get("ingredients")  # Expecting a list of ingredients

    if not user_id or not ingredients or not isinstance(ingredients, list):
        return jsonify({"message": "User ID and a list of ingredients are required"}), 400

    try:
        for ingredient in ingredients:
            query = """
            DELETE FROM shopping_list_items
            WHERE user_id = %s AND ingredient = %s
            """
            cursor.execute(query, (user_id, ingredient))

        db.commit()
        return jsonify({"message": f"{len(ingredients)} items removed from shopping list"}), 200
    except mysql.connector.Error as err:
        db.rollback()
        return jsonify({"message": f"Error: {err}"}), 500

if __name__ == "__main__":
    app.run(debug=True)
