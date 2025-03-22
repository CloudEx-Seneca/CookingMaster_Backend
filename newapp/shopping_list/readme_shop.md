# Shopping List
## Install dependences (mysql, flask, python)

```
python pyyaml
```
```
python mysql-connector-python
```
```
python Flask
```
## Initialize database

```
CREATE DATABASE IF NOT EXISTS shopping_list;
```
```
CREATE TABLE IF NOT EXISTS shopping_list_items (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    ingredient VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
```

## Run appllication

```
python shop_list.py
```

## Running functions
### Add Item

```
http://127.0.0.1:5000/add_item
```
![image](https://github.com/user-attachments/assets/255aa6d8-0714-4e96-aeb8-47f70243b0f8)

```
http://127.0.0.1:5000/get_list/<user_id>
```
![image](https://github.com/user-attachments/assets/7525d82a-598e-44ff-be48-c192e3c094ba)

```
http://127.0.0.1:5000/remove_item
```
![image](https://github.com/user-attachments/assets/035629cd-8338-4c9f-a9a4-39a643cafe75)

