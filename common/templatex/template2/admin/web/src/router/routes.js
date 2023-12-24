// 公共路由，无须鉴权
export const PublicRoutes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('src/pages/login/index.vue'),
    children: [],
  },
  {
    path: '/adminlogin',
    name: 'adminlogin',
    component: () => import('src/pages/login/index.vue'),
    children: [],
  },
  // {
  //   path: "/new-tab",
  //   component: () => import("layouts/NewTabLayout/index.vue"),
  //   children: [],
  // },

  // // Always leave this as last one,
  // // but you can also remove it

  // 以下内容在动态路由中添加，这里注释掉，解决刷新404的问题:store-permission-actions
  // {
  //   path: '/:catchAll(.*)*',
  //   name: 'notFound',
  //   component: () => import('pages/ErrorNotFound.vue'),
  // },
]

export const PrivateRoutes = [
  {
    path: '/',
    // name: "index",
    component: () => import('src/layouts/MainLayout/index.vue'),
    redirect: { path: '/dashboard' },
    children: [],
  },
]

export default [...PublicRoutes, ...PrivateRoutes]
