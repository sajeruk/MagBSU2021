import requests
import sys


def main():
    api_key = 'YOUR_API_KEY'
    n = int(sys.argv[1])
    min = int(sys.argv[2])
    max = int(sys.argv[3])


    url = "https://api.random.org/json-rpc/4/invoke"
    
    request  = {
        "jsonrpc": "2.0",
        "method": "generateIntegers",
        "params": {
            "apiKey": api_key,
            "n": n,
            "min": min,
            "max": max,
            "replacement": True
        },
        "id": 1
    }
    response = requests.post(url, json=request).json()
    print(response)

if __name__ == "__main__":
    main()
