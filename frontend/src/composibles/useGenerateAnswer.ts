import { createSharedComposable } from '@vueuse/core'
import { createWsClient } from '../fns/createWsClient'

export const useGenerateAnswer = createSharedComposable(() => {
  const answer = ref('')
  const isGeneratingAnswer = ref(false)
  const { addMessageHandler, sendDmMessage } = createWsClient({
    wsEndpoint: 'wss://aibridge.fly.dev/api/v1/ws',
    username: 'test',
    room: '79f7e192-f418-4b58-b75b-609761cd322b'
  })

  addMessageHandler<GenerateAnswerResponse>('generateAnswer/stream', (_sender: string, message: GenerateAnswerResponse) => {
    answer.value = message.message
  })
  addMessageHandler<GenerateAnswerResponse>('generateAnswer/done', (_sender: string, _message: GenerateAnswerResponse) => {
    isGeneratingAnswer.value = false
  })
  const generateAnswer = async (message: string) => {
    if (isGeneratingAnswer.value) return

    answer.value = ''
    isGeneratingAnswer.value = true
    sendDmMessage<GenerateAnswerRequest>('chatgpt', 'generateAnswer', {
      message: message,
      conversationID: 'latest'
    })
  }

  return {
    answer,
    isGeneratingAnswer,
    generateAnswer
  }
})
