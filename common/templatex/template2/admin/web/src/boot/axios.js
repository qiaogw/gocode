import { boot } from 'quasar/wrappers'
import axios from 'axios'
import { Notify, Dialog, Loading, QSpinnerGears } from 'quasar'
import { infoX, errorX } from './msg'
import { i18n } from './i18n'
import { useUserStore } from 'src/stores/user'
const userStore = useUserStore()
const api = axios.create({
  baseURL: process.env.BASE_URL,
  withCredentials: false,
})
export default boot(({ app, router, store }) => {
  const userStore = useUserStore()
  // console.log('app, router, storen', app, router, store)
  // 请求拦截
  api.interceptors.request.use(
    (request) => {
      // console.log('request', router, request)
      // fully customizable4
      Loading.show({
        spinner: QSpinnerGears,
        // other props
      })

      const token = userStore.GetToken()

      request.headers = {
        'Content-Type': 'application/json;charset=utf-8',
        'sub-Token': token,
        Authorization: token,
      }
      return request
    },
    (error) => {
      Loading.hide()
      errorX(error)
      return Promise.reject(error)
    }
  )
  // 响应拦截;
  api.interceptors.response.use(
    (response) => {
      Loading.hide()
      // const userStore = useUserStore()
      // // 如果JWT的ExpiresAt已经过期，但是RefreshAt没有过期，那么后台会在headers里插入sub-Refresh-Token，这里保存下来，形成更换token逻辑
      if (response.headers['sub-refresh-token']) {
        userStore.setToken(response.headers['sub-refresh-token'])
        // store.dispatch('user/SetToken', response.headers['sub-refresh-token'])
        Notify.create({
          type: 'positive',
          message:
            i18n.global.t('Refresh') + 'Token' + i18n.global.t('Success'),
        })
        return api(response.config)
      }
      const responseData = response.data
      const { code } = responseData
      if (code === undefined) {
        // 如果没有 code 代表这不是项目后端开发的接口 比如可能是  请求最新版本
        return response
      }
      if (code === 0) {
        if (!response.data.data) {
          infoX(response.data.msg)
          return response.data
        }
        return response.data.data
      } else {
        errorX(response.data.msg)
        // return response
      }
    },
    ({ response }) => {
      Loading.hide()
      let msg = '未知错误'
      let tmsg = true
      if (response) {
        msg = ''
        switch (response.status) {
          case -1:
            msg += '连接失败'
            break
          case 500:
            msg += '内部错误'
            break
          case 404:
            msg += '资源不存在'
            break
          case 401:
            tmsg = false
            Dialog.create({
              title: i18n.global.t('Authentication') + i18n.global.t('Failed'),
              message:
                response.data.message ||
                i18n.global.t('Please') + i18n.global.t('Relogin'),
              persistent: true,
              ok: {
                push: true,
                color: 'negative',
                label: i18n.global.t('Relogin'),
              },
            }).onOk(() => {
              // const userStore = useUserStore()
              userStore.HandleLogout()
              router.push({ name: 'login' })
            })
            break
          case 403:
            tmsg = false
            msg += '权限不足,请重新登录或联系管理员授权'
            Dialog.create({
              title: i18n.global.t('Authentication') + i18n.global.t('Failed'),
              message:
                response.data.message || '权限不足,请重新登录或联系管理员授权',
              persistent: true,
              ok: {
                push: true,
                color: 'negative',
                label: i18n.global.t('Relogin'),
              },
            }).onOk(() => {
              const userStore = useUserStore()
              userStore.HandleLogout()
              router.push({ name: 'login' })
            })
            errorX(msg)
            break
          default:
            msg += '未知错误'
        }
      } else {
      }
      if (tmsg) {
        errorX(msg)
      }
      return Promise.reject(response)
    },
    (error) => {
      Loading.hide()
      // 500的情况
      if (error + '' === 'Error: Request failed with status code 500') {
        Dialog.create({
          title: i18n.global.t('Error'),
          message:
            i18n.global.t('Data') +
            i18n.global.t('Exception') +
            ',' +
            i18n.global.t('Please') +
            i18n.global.t('Relogin'),
          persistent: true,
          ok: {
            push: true,
            color: 'negative',
            label: i18n.global.t('Logout'),
          },
        }).onOk(() => {
          // const userStore = useUserStore()
          userStore.HandleLogout()
          store.dispatch('user/HandleLogout')
          router.push({ name: 'login' })
        })
      }
      // 超时
      if (error + '' === 'Error: timeout of 40000ms exceeded') {
        Notify.create({
          type: 'negative',
          message: i18n.global.t('Operation') + i18n.global.t('Timeout'),
        })
      }
      // 网络错误情况，比如后台没有对应的接口
      if (error + '' === 'Error: Network Error') {
        router.push({ name: 'notFound' })
      } else if (error.response && error.response.status === 404) {
        // console.log(
        //   '请求地址不存在 [' + error.response.request.responseURL + ']'
        // )
        Notify.create({
          type: 'negative',
          message:
            i18n.global.t('Request') +
            i18n.global.t('Address') +
            i18n.global.t('NotFound') +
            ' ' +
            error.response.request.responseURL,
        })
      }
      return Promise.reject(error)
    }
  )

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export { api }
