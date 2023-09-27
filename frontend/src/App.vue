<script lang="ts" setup>
import Reader from './components/Reader.vue'
import ChatUI from './components/ChatUI.vue'
import { ref } from 'vue'

const selectedText = ref('')
const title = ref('')
</script>

<template>
  <ElConfigProvider>
    <ElContainer :class="$style.container">
      <ElContainer>
        <ElHeader :class="$style.header">Header</ElHeader>
        <ElMain :class="$style.main">
          <Reader v-model:selected-text="selectedText" v-model:title="title" />
        </ElMain>
      </ElContainer>
      <ElAside width="400px">
        <ChatUI :selected-text="selectedText" :title="title" />
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
  height: auto
  padding: 0
</style>
