import { defineStore } from 'pinia'
import { Cookies, SessionStorage, LocalStorage } from 'quasar'
import { login, getMe } from 'src/api/admin/auth'
// import { usePermissionStore } from 'src/stores/permission'
// import { useDictionaryStore } from 'src/stores/dict'
// const permissionStore = usePermissionStore()
// const dictionaryStore = useDictionaryStore()

export const useUserStore = defineStore('user', {
  state: () => ({
    info: undefined,
    token: undefined,
    xToken: undefined,
    rememberMe: true,
    language: undefined,
  }),

  getters: {},

  actions: {
    setInfo(v) {
      this.info = v
    },
    setToken(v) {
      this.token = v
      if (this.rememberMe) {
        LocalStorage.set('sub-token', v)
      } else {
        SessionStorage.set('sub-token', v)
        Cookies.set('sub-token', v)
      }
    },
    setXToken(v) {
      this.xToken = v
      if (this.rememberMe) {
        LocalStorage.set('sub-xToken', v)
      } else {
        SessionStorage.set('sub-xToken', v)
        Cookies.set('sub-xToken', v)
      }
    },
    setRememberMe(v) {
      this.rememberMe = v
    },
    setLanguage(v) {
      this.language = v
      LocalStorage.set('sub-language', v)
    },
    async isAdmin() {
      if (!this.info) {
        this.GetInfo()
      } else {
        return this.info.role.isAdmin
      }
    },
    GetToken() {
      if (this.token) {
        return this.token
      } else if (SessionStorage.getItem('sub-token')) {
        return SessionStorage.getItem('sub-token')
      } else if (LocalStorage.getItem('sub-token')) {
        return LocalStorage.getItem('sub-token')
      } else if (Cookies.get('sub-token')) {
        return Cookies.get('sub-token')
      }
    },
    GetXToken() {
      if (this.xToken) {
        return this.xToken
      } else if (SessionStorage.getItem('sub-xToken')) {
        return SessionStorage.getItem('sub-xToken')
      } else if (LocalStorage.getItem('sub-xToken')) {
        return LocalStorage.getItem('sub-xToken')
      } else if (Cookies.get('sub-xToken')) {
        return Cookies.get('sub-xToken')
      }
    },
    GetInfo() {
      if (this.info) {
        // console.log(this.info)
        return this.info
      } else {
        let info = this.getMeFetch()
        return info
      }
    },
    GetLanguage() {
      if (this.language) {
        return this.language
      } else if (LocalStorage.getItem('sub-language')) {
        return LocalStorage.getItem('sub-language')
      } else {
        return 'zh-CN'
      }
    },
    async Login(loginForm) {
      const res = await login(loginForm)
      if (res) {
        const token = res.accessToken
        const info = res.user
        this.setToken(token)
        this.setInfo(info)
        return true
      }
      return false
    },
    async getMeFetch() {
      const res = await getMe()
      if (res) {
        const info = res
        this.setInfo(info)
      }
      return res
    },
    clear() {
      delete SessionStorage.userInfo
      this.info = {}
      delete SessionStorage.token
      this.token = ''
      delete SessionStorage.xToken
      this.xToken = ''
    },
    HandleLogout() {
      // const permissionStore = usePermissionStore();
      // permissionStore.ClearMenu();
      SessionStorage.remove('sub-token')
      LocalStorage.remove('sub-token')
      Cookies.remove('sub-token')
      SessionStorage.remove('sub-xToken')
      SessionStorage.clear()
      LocalStorage.remove('sub-xToken')
      Cookies.remove('sub-xToken')
      // 字典不删除
      // LocalStorage.remove('sub-dict')
      this.token = undefined
      this.xToken = undefined
      this.info = undefined
      // permissionStore.ClearMenu()
      // dictionaryStore.clear()
    },
  },
})
