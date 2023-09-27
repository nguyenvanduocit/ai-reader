<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from 'vue'
import Epub, { Rendition, Contents } from 'epubjs'
import View from 'epubjs/types/managers/view'
import { PackagingMetadataObject } from 'epubjs/types/packaging'

const renderContainer = ref<HTMLDivElement>()
const book = Epub('/books/Machine Learning for Designers (Patrick Hebron) (Z-Library).epub')
let rendition: Rendition | undefined

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
const emits = defineEmits({
  'update:selectedText': (value: string) => true,
  'update:title': (value: string) => true
})

async function onSelected(cfiRange: string, contents: Contents) {
  if (!contents.window) {
    return
  }

  if (contents.window.getSelection()?.toString() === '') {
    emits('update:selectedText', '')
    return
  }

  const text = contents.window.getSelection()?.toString() || ''
  emits('update:selectedText', text)
}
onMounted(() => {
  book.loaded.metadata.then(function (meta: PackagingMetadataObject) {
    emits('update:title', meta.title)
  })
  rendition = book.renderTo(renderContainer.value!, {
    manager: 'continuous',
    flow: 'scrolled-doc',
    snap: true,
    width: '100%',
    allowScriptedContent: true
  })

  rendition.display(3)
  rendition.on('selected', onSelected)
  rendition.hooks.render.register(async ({ contents }: { contents: Contents }) => {
    await contents.addScript('https://cdn.jsdelivr.net/npm/darkreader@4.9.58/darkreader.min.js')
    contents.window?.DarkReader?.auto({
      brightness: 90,
      contrast: 80,
      sepia: 10
    })
  })
})

onBeforeUnmount(() => {
  rendition?.destroy()
  book.destroy()
})
</script>

<template>
  <div :class="$style.container" ref="renderContainer"></div>
</template>

<style module lang="sass">
.container
  width: 100%
  height: 100%
</style>
