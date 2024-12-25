const { kafka } = require("./client");
const readline = require("readline");

const TOPIC = "hashtag-topic"
const TOTAL_PARTITION = 3;


const allHashtag = [
  "InspirationDaily",
  "InspirationDaily",
  "InspirationDaily",
  "InspirationDaily",
  "CodeLife",
  "CodeLife",
  "CodeLife",
  "TechTrends",
  "TechTrends",
  "NatureLovers",
  "WeekendVibes",
  "HealthGoals",
  "ArtAndSoul",
  "TravelDiaries",
  "BookWorm",
  "MotivationMonday",
];

function findPartition(st, totalPartition) {
  let cnt = 0;
  for (const c of st) {
    cnt += c.charCodeAt(0); // Convert character to its ASCII value
    cnt %= totalPartition;
  }
  return cnt;
}

async function generateHashtag(producer) {
  const size = allHashtag.length;

  for (let i = 0; i < 10000; i++) {
    const postId = i;
    for (let j = 0; j < 3; j++) {
      const randInt = Math.floor(Math.random() * size);
      const hashTag = allHashtag[randInt];
      const hashtagPartition = findPartition(hashTag, TOTAL_PARTITION);

      await producer.send({
        topic: TOPIC,
        messages: [
          {
            partition: hashtagPartition,
            key: `${postId}`,
            value: JSON.stringify({ postId, hashTag, hashtagPartition,  }),
          },
        ],
      });

    }
  }
}

async function init() {
  const producer = kafka.producer();

  console.log("connecting producer");
  await producer.connect();
  console.log("producer connected");

  await generateHashtag(producer)
  await producer.disconnect();
}

init();
