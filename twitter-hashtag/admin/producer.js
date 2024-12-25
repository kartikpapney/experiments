const { kafka } = require("./client");
const readline = require("readline");

const TOPIC = "hashtag-topic"
const PARTITION = 3;

async function init() {
  const producer = kafka.producer();

  console.log("connecting producer");
  await producer.connect();
  console.log("producer connected");

  await producer.send({
    topic: "rider-updates",
    messages: [
      {
        partition: location.toLowerCase() === "north" ? 0 : 1,
        key: "location-update",
        value: JSON.stringify({ name: riderName, location }),
      },
    ],
  });

  await producer.disconnect();
}

init();
