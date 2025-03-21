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

# Function to delete an ingredient from the shopping list
@app.route("/remove_item", methods=["POST"])
def del_item(user_id, ingredient_id):
    query = "DELETE FROM shopping_list_items WHERE user_id = %s AND ingredient_id = %s"
    cursor.execute(query, (user_id, ingredient_id))
    db.commit()
    print("Ingredient removed from shopping list.")
    
if __name__ == "__main__":
    app.run(debug=True)
