# CLI test for Kafka end-to-end
# Bash script for local testing
# TODO: Replace for CI pipelines, scaffold in pure pytest

set -e # Exit on error

TOPIC="test-topic"
MESSAGE="hello-kafka-$(date +%s)" # Unique message
BROKER="localhost:9092"

echo "🔁 Producing message: $MESSAGE"
# Publishes message to kafka topic
echo "$MESSAGE" | kcat -P -b "$BROKER" -t "$TOPIC"

# Wait a moment for Kafka to process the message
sleep 1

echo "🔎 Consuming from topic: $TOPIC"
# Consumes the last message from the topic
OUTPUT=$(kcat -C -b "$BROKER" -t "$TOPIC" -e -q -o -1 -c 1)

echo "📨 Received: $OUTPUT"

if [[ "$OUTPUT" == "$MESSAGE" ]]; then
  echo "✅ Kafka test passed!"
  exit 0
else
  echo "❌ Kafka test failed."
  exit 1
fi
