<script setup lang="ts">
import { computed } from 'vue'
import { Mdit } from '../services/mdit'
import { useGenerateAnswer } from '../composibles/useGenerateAnswer'
const { answer, generateAnswer, isGeneratingAnswer } = useGenerateAnswer()

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

const outputHtml = computed(() => {
  return Mdit(answer.value)
})

const canSubmit = computed(() => {
  return props.selectedText.length > 0 && !isGeneratingAnswer.value
})

const translate = () => {
  generateAnswer('Act like the author of the book "' + props.title + '". Rewrite the following quote into Vietnamese: ' + props.selectedText)
}
</script>

<template>
  <div :class="$style.container">
    <el-progress :percentage="50" :show-text="false" :indeterminate="true" v-show="isGeneratingAnswer" />
    <ElButtonGroup :class="$style.controllers" v-show="!isGeneratingAnswer">
      <ElButton :disabled="!canSubmit" @click.prevent="translate">Translate</ElButton>
    </ElButtonGroup>
    <ElButtonGroup :class="$style.controllers" v-show="isGeneratingAnswer">
      <ElButton>cancel</ElButton>
    </ElButtonGroup>
    <ElText tag="div" :class="$style.output" v-html="outputHtml"></ElText>
  </div>
</template>

<style module lang="sass">
.container
  height: 100%
  padding: 10px
  :global(.el-progress--line)
    width: 100%
.controllers
  margin-bottom: 10px
  width: 100%

.output
  word-break: break-word

  :global(p)
    padding: 0
    margin: 0 0 25px
    line-height: 1.25rem
</style>
