import os
from flask import Flask, request, jsonify
from flask_cors import CORS
import mysql.connector
from mysql.connector import errorcode
import yaml

# Utility function to load database config from environment variables or fallback to YAML file
def load_db_config():
    db_config = {
        'host': os.getenv('DB_HOST'),
        'password': os.getenv('MYSQL_ROOT_PASSWORD')
    }

    return db_config

# Utility function to establish database connection
def create_db_connection(config):
    try:
        return mysql.connector.connect(
            host=config["host"],
            user="root",
            password=config["password"],
            database="shopping_list"
        )
    except mysql.connector.Error as err:
        handle_db_error(err)
        exit()

# Error handling for DB connection issues
def handle_db_error(err):
    error_messages = {
        errorcode.ER_ACCESS_DENIED_ERROR: "Something is wrong with your username or password",
        errorcode.ER_BAD_DB_ERROR: "Database does not exist"
    }
    print(error_messages.get(err.errno, err))
    
# Flask application setup
app = Flask(__name__)
CORS(app)  # Enable CORS

# Load config and initialize DB connection
config = load_db_config()
db = create_db_connection(config)

# Create a cursor before each request
@app.before_request
def create_cursor():
    global cursor
    cursor = db.cursor()

# Close the cursor after each request
@app.teardown_request
def close_cursor(exception=None):
    global cursor
    if cursor:
        cursor.close()

# Route to add items to the shopping list
@app.route("/shoppinglist/v1/add_item", methods=["POST"])
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

# Route to get items from the shopping list
@app.route("/shoppinglist/v1/get_list/<int:user_id>", methods=["GET"])
def get_shopping_list(user_id):
    try:
        query = "SELECT ingredient FROM shopping_list_items WHERE user_id = %s"
        cursor.execute(query, (user_id,))
        items = [item[0] for item in cursor.fetchall()]
        return jsonify(items), 200
    except mysql.connector.Error as err:
        return jsonify({"message": f"Error: {err}"}), 500

# Route to remove items from the shopping list
@app.route("/shoppinglist/v1/remove_item", methods=["POST"])
def remove_item():
    data = request.json
    user_id = data.get("user_id")
    ingredients = data.get("ingredients")

    if not user_id or not ingredients or not isinstance(ingredients, list):
        return jsonify({"message": "User ID and a list of ingredients are required"}), 400

    try:
        for ingredient in ingredients:
            query = "DELETE FROM shopping_list_items WHERE user_id = %s AND ingredient = %s"
            cursor.execute(query, (user_id, ingredient))

        db.commit()
        return jsonify({"message": f"{len(ingredients)} items removed from shopping list"}), 200
    except mysql.connector.Error as err:
        db.rollback()
        return jsonify({"message": f"Error: {err}"}), 500

# Run the app
if __name__ == "__main__":
    # Get the port from the environment variable or default to 5000
    port = int(os.getenv("SERVER_PORT", 5000))
    app.run(debug=True, port=port)
