import { boot } from 'quasar/wrappers'
import { Notify } from 'quasar'
const errorX = (msg, html) =>
  Notify.create({
    message: msg,
    timeout: 5000,
    html: !!html,
    color: 'negative',
  })
const infoX = (msg, html) =>
  Notify.create({
    message: msg,
    html: true,
    color: 'positive',
  })
const warnX = (msg, html) =>
  Notify.create({
    message: msg,
    html: !!html,
    color: 'warning',
  })
const retMsg = (r) => {
  if (r.success) {
    this.infoX(r.message)
  } else {
    this.errorX(r.message)
  }
}
// "async" is optional;
// more info on params: https://v2.quasar.dev/quasar-cli/boot-files
export default boot(({ app }) => {
  app.config.globalProperties.$error = errorX
  app.config.globalProperties.$info = infoX
  app.config.globalProperties.$warn = warnX
  app.config.globalProperties.$retMsg = retMsg
})

export { infoX, errorX, warnX, retMsg }
