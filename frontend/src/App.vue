<script lang="ts" setup>
import Epub, {Rendition} from 'epubjs'
import {onBeforeUnmount, onMounted, ref} from "vue";
import Section from "epubjs/types/section";
import Contents from "epubjs/types/contents";

const renderContainer = ref<HTMLDivElement>()
const book = Epub("/books/Machine Learning for Designers (Patrick Hebron) (Z-Library).epub")
let rendition: Rendition | undefined
const selectedText = ref<string>()

import { OpenAI } from "langchain/llms/openai";
import {ChatOpenAI} from "langchain/dist/chat_models/openai";

const llm = new OpenAI({
  openAIApiKey: "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",
});
const chatModel = new ChatOpenAI();

onMounted(() => {
  rendition = book.renderTo(renderContainer.value!, {
    snap: true,
    width: "100%",
    allowScriptedContent: true,
  });
  rendition.display(0);
  rendition.on("rendered", (section: Section) => {

  });

  rendition.on("relocated", (location: any) => {
    console.log(location)
  });

  rendition.on("selected", async function (cfiRange: string, contents: Contents) {
    if (!contents.window) {
      return
    }

    if (contents.window.getSelection()?.toString() === "") {
      selectedText.value = undefined
      return
    }


    const text = "Translate this text into Vietnamese: " + contents.window.getSelection()?.toString()

    const llmResult = await llm.predict(text);

    selectedText.value = llmResult
  });
})

const nextPage = () => {
  rendition?.next();
}

const previous = () => {
  rendition?.prev();
}

onBeforeUnmount(() => {
  rendition?.destroy()
  book.destroy()
})

</script>

<template>
  <div :class="$style.container">
    <div :class="$style.main">
      <div ref="renderContainer"></div>
    </div>
    <div :class="$style.sidebar">
      <button @click="previous">Previous</button>
      <button @click="nextPage">Next</button>
      <p>{{ selectedText }}</p>
    </div>
  </div>
</template>


<style lang="sass" module>
.container
  width: 100vw
  height: 100vh
  display: flex
  overflow: hidden

.main
  flex-grow: 1
  width: calc(100vw - 300px)
  overflow-y: auto

.sidebar
  width: 300px
  flex-grow: 0
  flex-shrink: 0
  border-left: 1px solid #ccc
  padding: 10px
</style>
