const { kafka } = require("./client");

const TOPIC = "hashtag-topic"
const PARTITION = 3;

async function init() {
  const admin = kafka.admin();
  console.log("admin connecting...");
  await admin.connect();
  console.log("admin connected...");

  console.log(`creating topic [${TOPIC}]`);
  await admin.createTopics({
    topics: [
      {
        topic: TOPIC,
        numPartitions: PARTITION,
      },
    ],
  });
  console.log(`topic created [${TOPIC}] with ${PARTITION} partitions`);
  console.log("disconnecting admin..");
  await admin.disconnect();
}

init();
