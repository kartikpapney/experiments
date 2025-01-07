import dotenv from "dotenv";
dotenv.config();
import { ChatOpenAI } from "@langchain/openai";
import { HumanMessage } from "@langchain/core/messages";
import { getSubtitles } from 'youtube-captions-scraper';

const model = new ChatOpenAI({ model: "gpt-4o-mini" });

function getPrompt(text) {
  return `

  Provide markdown-formatted notes for the lesson below. 
    - Focus on system design.
    - Be concise and avoid repetition.
    - Write for someone who is preparing for technical interview

    Text: "${text}"

`;
}


async function getDifficultWords() {
  try {
    
    const data = await getSubtitles({
      videoID: '5faMjKuB9bc', 
      lang: 'en'
    });

    const text = data.map((e) => e.text).join(" ");
 
    const response = await model.stream([
      new HumanMessage(getPrompt(text))
    ]);
    
    for await (const chunk of response) {
      process.stdout.write(chunk.content);
    }
  } catch(e) {
    console.log(e)
  }
}

getDifficultWords();