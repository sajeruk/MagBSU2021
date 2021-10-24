import json
import sys

def hello(name):
    return "hello " + name + "!"

def goodbye(name):
    return "goodbye " + name + "!"


def handle(request):
    obj = json.loads(request)
    method = globals()[obj['method']]
    params = obj['params']
    return method(**params)


if __name__ == "__main__":
    method = sys.argv[1]
    name = sys.argv[2]
    request = json.dumps({
        "method" : method,
        "params" : {"name" : name}
    })
    print("Request: " + request)
    print(handle(request))