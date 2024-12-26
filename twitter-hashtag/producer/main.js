const { kafka } = require("./client");
const readline = require("readline");

const TOPIC = "hashtag-topic"
const TOTAL_PARTITION = 3;


function findPartition(st, totalPartition) {
  let cnt = 0;
  for (const c of st) {
    cnt += c.charCodeAt(0); 
    cnt %= totalPartition;
  }
  return cnt;
}

async function generateHashtag(producer) {

  const POST_COUNT = parseInt(process.argv[2]);

  for (let i = 0; i < POST_COUNT; i++) {
    const postId = i;
    for (let j = 0; j < 3; j++) {
      
      const hashTag = process.argv[3];
      const hashtagPartition = findPartition(hashTag, TOTAL_PARTITION);

      await producer.send({
        topic: TOPIC,
        messages: [
          {
            partition: hashtagPartition,
            key: `${postId}`,
            value: JSON.stringify({ postId, hashTag, hashtagPartition }),
          },
        ],
      });

      console.log(`Post = ${postId} with hashtag ${hashTag} in partition ${hashtagPartition}`)
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
