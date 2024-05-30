from flask import Flask, jsonify, render_template

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('index.html')

@app.route('/api/data')
def api_data():
    data = {"message": "Hello from Flask API!"}
    return jsonify(data)

if __name__ == '__main__':
    app.run(debug=True)
