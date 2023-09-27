import { ChatGooglePaLM } from 'langchain/chat_models/googlepalm'

// To enable streaming, we pass in `streaming: true` to the LLM constructor.
// Additionally, we pass in a handler for the `handleLLMNewToken` event.
export const llm = new ChatGooglePaLM({
  apiKey: import.meta.env.VITE_GOOGLE_PALM_API_KEY,
  temperature: 0.7, // OPTIONAL
  modelName: 'models/chat-bison-001', // OPTIONAL
  topK: 40, // OPTIONAL
  topP: 3
})
