import subprocess
import time


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
