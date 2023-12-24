import { computed, onMounted, ref } from 'vue'
import { FormatDateTime } from 'src/utils/date'
import { useI18n } from 'vue-i18n'

export default function useCommon() {
  const subLogo = () => {
    console.info('欢迎使用sub!')
  }
  const { t } = useI18n

  const openLink = (url) => {
    window.open(url)
  }
  const showDateTime = computed(() => {
    return (datetime) => {
      return FormatDateTime(datetime)
    }
  })
  // 首页等允许无token的白名单
  const AllowList = ['/login']
  // 没有用户名的时候使用这个名字
  const subDefaultUsername = () => 'SUB用户'
  // 没有头像配置的时候使用这个头像
  const subDefaultAvatar = 'favicon-128x128.png'
  // 没有网站前台配置的时候用这个配置
  const subFrontendDefault = {
    mainTitle: 'Gin-Quasar-Admin',
    subTitle: 'Gin-Quasar-Admin',
    webDescribe: 'Be the change you want to see in the world.',
    showGit: 'yes',
  }
  const selectOptionLabel = (opt) => {
    if (
      opt.name === 'system' ||
      opt.parent_code === 'system' ||
      opt.parent_code === 'log'
    ) {
      return t(opt.title)
    }
    return opt.title
  }
  const selectRouteLabel = (opt) => {
    if (
      opt.name === 'system' ||
      opt.meta.parent_code === 'system' ||
      opt.meta.parent_code === 'log'
    ) {
      return t(opt.meta.title)
    }
    return opt.meta.title
  }
  // 没有网站后台配置的时候用这个配置
  const subBackendDefault = {}
  return {
    subLogo,
    showDateTime,
    openLink,
    AllowList,
    subDefaultUsername,
    subDefaultAvatar,
    subFrontendDefault,
    subBackendDefault,
    selectOptionLabel,
    selectRouteLabel,
  }
}
