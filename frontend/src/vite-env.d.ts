/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

interface Window {
  DarkReader: any
}

declare type WebsocketMessage<T> = T & {
  sender: string
  recipient: string
  type: string
}

interface BookHydrate {
  cover_url: string
  book_url: string
  hash: string
  progress: number
  title: string
  author: string
  modified: number
}

declare type MessageHandler = (sender: string, message: any) => void

declare type WsClientOptions = {
  wsEndpoint: string
  username: string
  room: string
}

declare type GenerateAnswerResponse = {
  message?: any
  conversationID?: string
  messageID?: string
}

declare type GenerateAnswerRequest = {
  message?: any
  conversationID?: string | null
  parentMessageID?: string | null
}
