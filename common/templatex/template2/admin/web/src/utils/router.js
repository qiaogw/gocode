import { PrivateRoutes } from 'src/router/routes'

export const HandleRouter = (menuData) => {
  const result = []
  for (let item of menuData) {
    if (item.path !== '') {
      // console.log("ment:", item);
      // console.log(pageImporter(item.component));
      const obj = {
        path: item.path,
        name: item.name,
        component: pageImporter(item.component),
        meta: {
          hidden: item.hidden,
          keepAlive: item.keepAlive,
          title: item.title,
          icon: item.icon,
          name: item.name,
          btn: item.button || undefined,
          remark: item.remark,
        },
        redirect: item.redirect,
      }
      result.push(obj)
    } else {
      if (item.is_link === 'yes') {
        delete item.path
      }
    }
  }
  // 将需要鉴权路由的children（默认是空的）替换成后台传过来的整理后的路由表
  PrivateRoutes[0].children = [...result]
  // 返回鉴权路由
  return PrivateRoutes
}

//**为通配符,vite不支持require导入方式,故用import.meta.glob(vite动态导入)
/*import.meta.glob
 * 该方法匹配到的文件默认是懒加载，通过动态导入实现，构建时会分离独立的 chunk，是异步导入，返回的是 Promise
 * /*import.meta.globEager
 * 该方法是直接导入所有模块，并且是同步导入，返回结果直接通过 for...in循环就可以操作
 * 后端返回来的路由component是字符串,如component: "pages/Index/index.vue",
 * 前端需要把component: "pages/Index/index.vue" 转化为组件对象
 * component:() => import("/src/pages/Index/index.vue")
 **/
const routeAllPathToCompMap = import.meta.glob(`../pages/**/*.vue`)

const pageImporter = (component) => {
  // Quasar2 版本: src\pages\dashboard\index.vue
  return routeAllPathToCompMap[`../${component}.vue`]
  // Quasar1 版本:
  // return (resolve) => require([`src/pages/${component}`], resolve)
}
