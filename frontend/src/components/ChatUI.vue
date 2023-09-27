<script setup lang="ts">
import { ref, computed } from 'vue'
import { llm } from '../services/llm'
import { LLMResult } from 'langchain/schema'
import { Serialized } from 'langchain/dist/load/serializable'
import { Mdit } from '../services/mdit'
const isStreaming = ref(false)
const props = defineProps({
  selectedText: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  }
})

const output = ref('')
const outputHtml = computed(() => {
  return Mdit(output.value)
})

const canSubmit = computed(() => {
  return props.selectedText.length > 0
})

const callPredict = (prompt: string) => {
  console.log(prompt)
  output.value = 'Translating...'
  llm.predict(prompt, {
    callbacks: [
      {
        handleLLMStart(
          llm: Serialized,
          prompts: string[],
          runId: string,
          parentRunId?: string,
          extraParams?: Record<string, unknown>,
          tags?: string[],
          metadata?: Record<string, unknown>
        ): Promise<void> | void {
          output.value = ''
          isStreaming.value = true
        },
        handleLLMNewToken(token: string) {
          output.value += token
        },
        handleLLMError(err: Error | unknown) {
          if (err instanceof Error) {
            output.value = err.message
          } else {
            output.value = 'Unknown error'
          }
        },
        handleLLMEnd(output: LLMResult, runId: string, parentRunId?: string, tags?: string[]): Promise<void> | void {
          isStreaming.value = false
        }
      }
    ]
  })
}

const translate = () => {
  callPredict('Act like the author of the book "' + props.title + '". Rewrite the following quote into Vietnamese: ' + props.selectedText)
}

const explain = () => {
  callPredict('Act like the author of the book "' + props.title + '". explain the following quote: ' + props.selectedText)
}

const giveOpinion = () => {
  callPredict('Act like the author of the book "' + props.title + '". Let\'s give some extended thoughts around this quote: ' + props.selectedText)
}

const summary = () => {
  callPredict('Act like the author of the book "' + props.title + '". Summarize the following quote, make it shorter, easier to understand: ' + props.selectedText)
}
</script>

<template>
  <div :class="$style.container">
    <ElButtonGroup v-loading="isStreaming" :class="$style.controllers">
      <ElButton :disabled="!canSubmit" @click.prevent="translate">Translate</ElButton>
      <ElButton :disabled="!canSubmit" @click.prevent="explain">Explain</ElButton>
      <ElButton :disabled="!canSubmit" @click.prevent="summary">Summary</ElButton>
      <ElButton :disabled="!canSubmit" @click.prevent="giveOpinion">Deep dive</ElButton>
    </ElButtonGroup>
    <ElText tag="div" :class="$style.output" v-html="outputHtml"></ElText>
  </div>
</template>

<style module lang="sass">
.container
  height: 100%
  padding: 10px
.controllers
  margin-bottom: 10px
.output
  word-break: break-word
  :global(p)
    padding: 0
    margin: 0 0 25px
    line-height: 1.25rem
</style>
