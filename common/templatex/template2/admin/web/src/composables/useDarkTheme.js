import { useQuasar } from 'quasar'
import { computed } from 'vue'

export default function useDarkTheme() {
  const $q = useQuasar()
  const darkTheme = computed(() => {
    if ($q.dark.isActive) {
      return 'bg-dark text-white'
    } else {
      return 'bg-primary '
    }
  })
  const darkThemeSelect = computed(() => {
    if ($q.dark.isActive) {
      return 'bg-grey-9 text-orange'
    } else {
      return 'bg-blue-1 '
    }
  })
  const darkThemeChart = computed(() => {
    if ($q.dark.isActive) {
      return 'dark'
    } else {
      return ''
    }
  })
  return {
    darkTheme,
    darkThemeSelect,
    darkThemeChart,
  }
}
