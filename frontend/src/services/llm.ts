import { ChatOpenAI } from 'langchain/chat_models/openai'

// To enable streaming, we pass in `streaming: true` to the LLM constructor.
// Additionally, we pass in a handler for the `handleLLMNewToken` event.
export const llm = new ChatOpenAI({
  maxTokens: 2000,
  streaming: true,
  openAIApiKey: import.meta.env.VITE_OPENAI_KEY,
  cache: true
})
