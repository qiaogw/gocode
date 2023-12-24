import comSearch from 'src/components/comSearch/index.vue'
import comUpload from 'src/components/comUpload/index.vue'
import comDel from 'src/components/comdel/index.vue'
import comEdit from 'src/components/comEdit/index.vue'
import comCode from 'src/components/comCode/index.vue'
import copyButton from 'src/components/copyButton/index.vue'

export default ({ app }) => {
  // console.log("props");
  app.component('com-del', comDel)
  app.component('com-edit', comEdit)
  app.component('com-search', comSearch)
  app.component('com-upload', comUpload)
  app.component('com-code', comCode)
  app.component('CopyButton', copyButton)
}
