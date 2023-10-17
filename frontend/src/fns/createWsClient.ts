import ReconnectingWebSocket, { Options } from 'reconnecting-websocket'
import * as Events from 'reconnecting-websocket/events'

const messageHandlers: Record<string, MessageHandler> = {}

export const createWsClient = (options: WsClientOptions) => {
  const websocketOptions: Options = {
    connectionTimeout: 1000,
    startClosed: false
  }

  const ws = new ReconnectingWebSocket(`${options.wsEndpoint}?room=${options.room}&username=${options.username}`, [], websocketOptions)

  ws.addEventListener('message', (event) => {
    const message = JSON.parse(event.data) as WebsocketMessage<any>
    if (message.type === undefined) {
      return
    }

    const messageType = message.type
    if (messageHandlers[messageType] !== undefined) {
      messageHandlers[messageType].call(null, message.sender, message)
    }
  })

  const addEventListener = <T extends keyof Events.WebSocketEventListenerMap>(type: T, listener: Events.WebSocketEventListenerMap[T]) => {
    ws.addEventListener(type, listener)
  }

  const sendDmMessage = <T>(recipient: string, type: string, message: T) => {
    const wsMessage: WebsocketMessage<T> = {
      type: type,
      sender: options.username,
      recipient: recipient,
      ...message
    }

    ws.send(JSON.stringify(wsMessage))
  }

  const addMessageHandler = <T>(type: string, handler: (sender: string, message: T) => void) => {
    if (messageHandlers[type.toString()] !== undefined) {
      throw new Error(`Message handler for type ${type.toString()} already exists`)
    }

    messageHandlers[type.toString()] = handler
  }

  const removeMessageHandler = (type: string) => {
    if (messageHandlers[type.toString()] === undefined) {
      throw new Error(`Message handler for type ${type.toString()} does not exist`)
    }

    delete messageHandlers[type]
  }

  const removeAllMessageHandlers = () => {
    Object.keys(messageHandlers).forEach((key) => {
      removeMessageHandler(key)
    })
  }

  const open = () => {
    ws.reconnect()
  }

  const close = () => {
    removeAllMessageHandlers()
    ws.close()
  }

  return {
    open,
    close,
    addEventListener,
    addMessageHandler,
    removeMessageHandler,
    sendDmMessage
  }
}
