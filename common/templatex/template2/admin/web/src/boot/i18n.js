import { boot } from 'quasar/wrappers'
import { createI18n } from 'vue-i18n'
import messages from 'src/i18n'
// import { useUserStore } from 'src/stores/user'

// export const userStore = useUserStore()
const i18n = createI18n({
  locale: 'zh-CN',
  globalInjection: true,
  fallbackLocale: 'zh-CN',
  messages,
  silentTranslationWarn: true,
  silentFallbackWarn: true,
})
export default boot(({ app }) => {
  // Set i18n instance on app
  app.use(i18n)
})

export { i18n }
