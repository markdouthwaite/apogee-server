from apogee.models import load_model
from apogee.server import ApogeeServer


def run(port=8080, host="127.0.0.1"):
    model = load_model("alarm")
    server = ApogeeServer(model, port=port, host=host)
    server.run()


if __name__ == "__main__":
    run()
