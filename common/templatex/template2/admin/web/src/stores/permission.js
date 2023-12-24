import { defineStore } from 'pinia'
import { getMenu } from 'src/api/admin/auth'
import { HandleRouter } from 'src/utils/router'
import { ArrayToTree, TreeToArray, Array2Tree } from 'src/utils/arrayAndTree'

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    userMenu: [],
    searchMenu: [],
    topMenu: [],
  }),
  getters: {},
  actions: {
    async GetUserMenu() {
      const res = await getMenu()
      if (res) {
        // if (res.code === '100006') {
        //   console.log('100006')
        // }
        const data = res.list
        // 拿到鉴权路由表（用户自己的所有菜单），整理成路由
        const userMenu = HandleRouter(data)
        // 加入404界面
        userMenu.push({
          path: '/:catchAll(.*)*',
          name: 'notFound',
          component: () => import('src/pages/ErrorNotFound.vue'),
        })
        // 设置所有菜单
        this.InitUserMenu(userMenu)
        // 去掉隐藏菜单
        const searchMenu = data.filter((value) => !value.hidden)

        // 设置搜索菜单
        this.InitSearchMenu(searchMenu)
        // 深度拷贝，避免影响其他数据
        const searchMenuNew = JSON.parse(JSON.stringify(searchMenu))
        // 生成菜单树

        const topMenu = ArrayToTree(searchMenuNew, 'id', 'parentId')

        this.InitTopMenu(topMenu)
        // 返回鉴权路由表
        return userMenu
      } else {
        return
      }
    },
    InitUserMenu(routes) {
      const menu = []
      const push = function (routes) {
        routes.forEach((route) => {
          if (route.children && route.children.length > 0) {
            push(route.children)
          } else {
            if (route.meta && route.meta.btn && route.meta.btn.length > 0) {
              push(route.meta.btn)
            } else {
              const { meta, name, path } = route
              menu.push({ meta, name, path })
            }
          }
        })
      }
      push(routes)
      this.userMenu = menu
    },
    InitSearchMenu(searchMenu) {
      this.searchMenu = searchMenu
    },
    InitTopMenu(topMenu) {
      this.topMenu = topMenu
    },
    ClearMenu() {
      this.userMenu = []
      this.searchMenu = []
      this.topMenu = []
    },
  },
})
