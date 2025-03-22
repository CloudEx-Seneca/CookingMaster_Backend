from database import get_db_connection
import json

def add_to_shopping_list(data):
    db = get_db_connection()
    cursor = db.cursor()

    user_id = data["user_id"]
    ingredient_id = data["ingredient_id"]
    ingredient = data["ingredient"]

    query = """
    INSERT INTO shopping_list_items (user_id, ingredient_id, ingredient)
    VALUES (%s, %s, %s)
    ON DUPLICATE KEY UPDATE ingredient = VALUES(ingredient)
    """
    cursor.execute(query, (user_id, ingredient_id, ingredient))
    db.commit()

    cursor.close()
    db.close()

    return {"message": "Item added to shopping list"}
