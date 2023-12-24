import { boot } from 'quasar/wrappers'
import { LoadingBar, Loading, QSpinnerGears } from 'quasar'
import XEUtils from 'xe-utils'

import { useUserStore } from 'src/stores/user'
import { usePermissionStore } from 'src/stores/permission'

import useCommon from 'src/composables/useCommon'
import { i18n } from './i18n'
const userStore = useUserStore()
const permissionStore = usePermissionStore()

LoadingBar.setDefaults({
  color: 'red',
  size: '4px',
  position: 'top',
})

function startLoading() {
  Loading.show({
    // spinner: QSpinnerGears,
    message: i18n.global.t('System') + i18n.global.t('Loading'),
  })
  LoadingBar.start()
}

function stopLoading() {
  Loading.hide()
  LoadingBar.stop()
}

export default boot(({ app, router, store }) => {
  app.directive('permission', (el, binding, vNode, prevNode) => {
    const needPermissions = binding.value
    const permissions = permissionStore.userMenu
    // const permissions = ['function_edit', 'function_add']
    // const isAdmin = userStore.isAdmin();
    // if (isAdmin) {
    //   return true;
    // }
    const hasPermission = permissions.some((s) => {
      return (
        XEUtils.findKey(permissions, (item) => item.name === needPermissions) >
        -1
      )
    })
    const waitUse = needPermissions.toString().split(',')

    let flag = waitUse.some((item) => {
      return item === permissions
    })
    if (!hasPermission) {
      el.style.display = 'none'
    }
    // storageStore.SetSubDict()
    // 注意 prevNode 只有在 updated 生命周期才有值！
  })

  router.beforeEach((to, from, next) => {
    // console.log('beforeEach((', to, from, next)
    startLoading()
    const token = userStore.GetToken()
    const { AllowList } = useCommon()
    if (token) {
      if (AllowList.indexOf(to.path) !== -1) {
        next({ path: '/' })
        stopLoading()
      } else {
        if (!permissionStore.userMenu.length) {
          permissionStore.GetUserMenu().then((res) => {
            // 在vue-router4中，addRoutes被废弃，改为了addRoute，循环调用
            // 动态添加鉴权路由表
            if (res) {
              res.forEach((item) => {
                router.addRoute(item)
              })
              next({ ...to, replace: true })
            } else {
              userStore.HandleLogout()
              // store.dispatch('user/HandleLogout')
              next({ path: '/', replace: true })
            }
          })
        } else {
          next()
        }
        stopLoading()
      }
    } else {
      if (AllowList.indexOf(to.path) !== -1) {
        next()
        stopLoading()
      } else {
        next(`/login?redirect=${to.fullPath}`)
        stopLoading()
      }
    }
  })
  router.afterEach(() => {
    stopLoading()
  })
})
