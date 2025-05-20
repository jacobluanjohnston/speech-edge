import subprocess
import time
import socket
import pytest


def is_kafka_running(host="localhost", port=9092):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        return s.connect_ex((host, port)) == 0


@pytest.mark.skipif(not is_kafka_running(), reason="Kafka is not running locally")
def test_kafka_produce_consume():
    topic = "test-topic"
    message = f"hello-kafka-{int(time.time())}"

    # Send message
    subprocess.run(
        ["kcat", "-P", "-b", "localhost:9092", "-t", topic],
        input=message.encode(),
        check=True,
    )

    # Wait and consume
    time.sleep(1)
    result = subprocess.run(
        [
            "kcat",
            "-C",
            "-b",
            "localhost:9092",
            "-t",
            topic,
            "-e",
            "-q",
            "-o",
            "-1",
            "-c",
            "1",
        ],
        capture_output=True,
        check=True,
    )

    output = result.stdout.decode().strip()
    assert output == message, f"Expected {message}, got {output}"
