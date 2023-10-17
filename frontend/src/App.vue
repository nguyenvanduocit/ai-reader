<script setup lang="ts">
import Reader from './components/Reader.vue'
import ChatUI from './components/ChatUI.vue'
import { refDebounced } from '@vueuse/core'
import { ImportBook, OpenEpubDialog } from '../wailsjs/go/main/App'

const selectedText = ref('')
const debouncedSelectedText = refDebounced(selectedText, 2000)
const title = ref('')
const selectedUrl = ref('')
const isBookOpened = computed(() => selectedUrl.value !== '')
const bookId = ref('')
const openBook = async () => {
  const filePath = await OpenEpubDialog()
  try {
    const reponse = await ImportBook(filePath)
    console.log(reponse)
  } catch (e) {
    console.error('Failed to import:', e)
  }
}
</script>

<template>
  <ElConfigProvider>
    <ElContainer :class="$style.container">
      <ElContainer>
        <ElMain :class="$style.main">
          <Reader :key="bookId" :path="selectedUrl" v-model:selected-text="selectedText" />
        </ElMain>
      </ElContainer>
      <ElAside width="400px">
        <ElHeader :class="$style.header"><ElButton @click.prevent="openBook">Open book</ElButton></ElHeader>
        <ChatUI :selected-text="debouncedSelectedText" :title="title" />
      </ElAside>
    </ElContainer>
  </ElConfigProvider>
</template>

<style lang="sass" module>
.container
  width: 100vw
  height: 100vh
  overflow: hidden
.header
  height: 50px
  display: flex
  justify-content: center
  align-items: center
.main
  width: 100%
  height: 100vh
  padding: 0
</style>
