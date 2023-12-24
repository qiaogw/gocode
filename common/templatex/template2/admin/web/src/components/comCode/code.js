import 'prismjs'
import { h, computed } from 'vue'

export default {
  name: 'CodePrism',

  props: {
    code: String,
    lang: String,
  },

  setup(props) {
    const className = computed(() => `language-${props.lang}`)

    return () =>
      h(
        'pre',
        { class: ' ' + className.value, style: 'background-color: #1d1f21;' },
        [
          h('code', {
            class: 'line-numbers prism  has-numbering',

            innerHTML: Prism.highlight(
              props.code,
              Prism.languages[props.lang],
              props.lang
            ),
          }),
        ]
      )
  },
}
