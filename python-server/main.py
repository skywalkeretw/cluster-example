from os import environ
from flask import Flask, jsonify, request

app=Flask(__name__)

@app.route("/",methods=['GET','POST'])
def index():
    if request.method=='GET':
        name = "John Doe"     
        result = {"name": name}
        return jsonify(result)
    else:
        return jsonify({'Error':"This is a GET API method"})


if __name__ == '__main__':
    debug = False
    if environ.get('DEBUG') is not None and environ.get('DEBUG') == "True":
        debug = True

    app.run(debug=debug, host='0.0.0.0', port=8080)