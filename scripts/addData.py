import yaml
import psycopg2
import json
import base64
from csv import DictReader

def get_db_conn():
    conn = psycopg2.connect(
        host="localhost",
        # database="test",
        user="postgres", # Replace postgres user name
        password="Allhailbruno18") # Postgres password
    return conn

def cleanup_all_tables(db_conn, table_list):
    print("Cleaning up: ", ','.join(table_list))
    cursor = db_conn.cursor()
    for table in table_list:
        query = "TRUNCATE TABLE {table} CASCADE".format(table = table)
        print(query)
        cursor.execute(query)
        db_conn.commit()


def add_data_to_database(record, table_name):
    db_conn = get_db_conn()
    cursor = db_conn.cursor()
    column_names = []
    column_values = []
    # record = base64_encoding_flowUI(record)
    column_names.extend(record.keys())
    column_values.extend(record.values())
    # print("Column Names: ", column_names)
    # print("Column Values: ", column_values)             
    query = "INSERT INTO {table_name}({column_names_str}) VALUES ({column_values_str});".format(
        table_name=table_name,
        column_names_str=','.join(column_names),
        column_values_str=','.join(repr(value) for value in column_values)
    )
    print(query)
    cursor.execute(query)
    db_conn.commit()
                
def read_csv():
    table_name = 'subjects'
    with open('./scripts/'+table_name+'.csv', 'r', encoding='utf-8-sig') as file:
        dict_reader = DictReader(file)
        records = list(dict_reader)
        for record in records:
            record = {key:val.strip() for key, val in record.items()}
            # print(record)
            add_data_to_database(record, table_name)


if __name__ == "__main__":
    db_conn = get_db_conn()
    cleanup_all_tables(db_conn, ["subjects"])
    read_csv()