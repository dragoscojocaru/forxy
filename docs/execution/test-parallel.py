import requests
import json
import concurrent.futures
import time

with open("forxy-payload.json", 'r') as file:
    payload = json.loads(file.read())

def send_request(req_id, req_data):
    try:
        response = requests.get(req_data["url"], json=req_data["body"])
        print(f"Request {req_id}: Status code: {response.status_code}, Response: {response.text}")
    except Exception as e:
        print(f"Request {req_id} failed: {e}")

start = time.time()

with concurrent.futures.ThreadPoolExecutor() as executor:
    futures = [executor.submit(send_request, req_id, req_data) for req_id, req_data in payload['requests'].items()]

    for future in concurrent.futures.as_completed(futures):
        pass

end = time.time()
print("Execution time: " + str(end-start))
