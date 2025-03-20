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

# Function to display a user's shopping list
@app.route("/get_list/<int:user_id>", methods=["GET"])
def get_shopping_list(user_id):
    query = """
    SELECT i.name FROM shopping_list_items s
    JOIN recipe.ingredients i ON s.ingredient_id = i.id
    WHERE s.user_id = %s
    """
    cursor.execute(query, (user_id,))
    items = cursor.fetchall()
    for item in items:
        print(f"{item[0]}")

if __name__ == "__main__":
    app.run(debug=True)
