# filepath: /Users/eduardo/platzi/python/copilot-course/PythonAPI/app.py

from flask import Flask
from flasgger import Swagger # Importar Swagger

app = Flask(__name__)
swagger = Swagger(app) # Inicializar Swagger con la app

@app.route('/hello/<string:cadenaDeEntrada>', methods=['GET'])
def hello(cadenaDeEntrada):
    """
    Este endpoint recibe una cadena y devuelve un saludo personalizado.
    ---
    parameters:
      - name: cadenaDeEntrada
        in: path
        type: string
        required: true
        description: La cadena que se incluirá en el saludo.
    responses:
      200:
        description: Un saludo personalizado.
        schema:
          type: string
          example: Hola Mundo desde la API de Python
    """
    return f"Hola {cadenaDeEntrada} desde la API de Python"

if __name__ == '__main__':
    app.run(debug=True)