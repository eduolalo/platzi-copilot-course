# filepath: /Users/eduardo/platzi/python/copilot-course/PythonAPI/test_app.py
import pytest
from app import app as flask_app # Importa tu instancia de Flask app

@pytest.fixture
def app():
    """Crea y configura una nueva instancia de la app para cada prueba."""
    # Configuración opcional, por ejemplo, para usar una base de datos de prueba
    # flask_app.config.update({
    #     "TESTING": True,
    # })
    yield flask_app

@pytest.fixture
def client(app):
    """Un cliente de prueba para la app."""
    return app.test_client()

def test_hello_endpoint(client):
    """Prueba el endpoint /hello/<cadenaDeEntrada>."""
    test_string = "Prueba"
    response = client.get(f'/hello/{test_string}')

    # Verifica que el código de estado sea 200 OK
    assert response.status_code == 200

    # Verifica que el contenido de la respuesta sea el esperado
    expected_data = f"Hola {test_string} desde la API de Python"
    assert response.data.decode('utf-8') == expected_data

    # Prueba con otra cadena
    test_string_2 = "Mundo"
    response_2 = client.get(f'/hello/{test_string_2}')
    assert response_2.status_code == 200
    expected_data_2 = f"Hola {test_string_2} desde la API de Python"
    assert response_2.data.decode('utf-8') == expected_data_2