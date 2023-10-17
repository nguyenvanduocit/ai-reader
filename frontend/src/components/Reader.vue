<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from 'vue'

const renderContainer = ref<HTMLDivElement>()

const props = defineProps({
  selectedText: {
    type: String,
    default: ''
  },
  path: {
    type: String,
    default: ''
  }
})

const emits = defineEmits({
  'update:selectedText': (_value: string) => true,
  'update:title': (_value: string) => true
})

onBeforeUnmount(() => {})
</script>

<template>
  <div :class="$style.container" ref="renderContainer">
    <iframe
      width="100%"
      frameborder="0"
      referrerpolicy="unsafe-url"
      sandbox="allow-scripts"
      src="http://0.0.0.0:8080/Dark%20Psychology%20%28James%20W.%20Williams%29%20%28Z-Library%29/text/part0000_split_023.html"
    />
  </div>
</template>

<style module lang="sass">
.container
  width: 100%
  height: 100%
</style>
