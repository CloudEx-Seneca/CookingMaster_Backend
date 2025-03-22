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

![image](https://github.com/user-attachments/assets/255aa6d8-0714-4e96-aeb8-47f70243b0f8)
