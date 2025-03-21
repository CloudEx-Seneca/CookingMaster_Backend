Use readme.md in newapp!!! 
# For Shopping List
# Install dependencies
pip install flask mysql-connector-python pyyaml

# Import SQL schema
mysql -u root -p < shopping_list.sql

# Start flask server for each section you wish to use
python add_item.py
python display_list.py
python del_item.py

# Add for test