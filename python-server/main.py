from os import environ
from flask import Flask, jsonify


app=Flask(__name__)


@app.route("/")
def index():
    result = {"body": "Hello form Python"}
    return jsonify(result)

if __name__ == '__main__':
    debug = False
    if environ.get('DEBUG') is not None and environ.get('DEBUG') == "True":
        debug = True

    app.run(debug=debug, host='0.0.0.0', port=8080)