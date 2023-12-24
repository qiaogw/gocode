<template>
  <div class="doc-code-container">
    <code-prism
      :lang="lang"
      :code="code"
      class="line-numbers com-code relative-position"
    ></code-prism>
    <div class="doc-code-container__copy absolute">
      <copy-button :text="code"></copy-button>
      <!-- {{code}} -->
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import CodePrism from './code.js'
import 'prismjs/themes/prism.css'
// import CopyButton from './CopyButton.vue'
export default {
  name: 'DocCode',
  components: {
    CodePrism,
    // CopyButton
  },
  props: {
    code: String,
    maxHeight: String,
    lang: {
      type: String,
      default: 'js',
    },
  },
  setup(props) {
    const style = computed(() =>
      props.maxHeight !== void 0
        ? { overflow: 'auto', maxHeight: props.maxHeight }
        : null
    )
    return { style }
  },
}
</script>

<style lang="sass">
.doc-code-container
  &__copy
    top: 22px
    right: 16px
</style>
