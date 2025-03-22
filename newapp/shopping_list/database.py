import mysql.connector
import yaml

def get_db_connection():
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
    return db
