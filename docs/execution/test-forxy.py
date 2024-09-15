import requests
import time

with open("forxy-payload.json", 'r') as file:
    payload = file.read()

url = "http://localhost:1480/http/fork"

start = time.time()

response = requests.get(url, data=payload)

print(f"Status code: {response.status_code} Response body: {response.text}")

end = time.time()
print("Execution time: " + str(end-start))
