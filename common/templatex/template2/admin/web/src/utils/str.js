import XEUtils from 'xe-utils'
//搜索字符串
export const searchReg = function (name1, name2) {
  var patt = new RegExp(name2, 'i')
  return patt.test(name1)
}

//将带驼峰字符串转成蛇形字符串
export const snakeCase = function (name) {
  if (!name) {
    return ''
  }
  return name.replace(/([A-Z])/g, '_$1').toLowerCase()
}

//将蛇形字符串转成驼峰字符串
export const camelCase = function (name) {
  if (!name) {
    return ''
  }
  return name.replace(/\_(\w)/g, function (all, letter) {
    return letter.toUpperCase()
  })
}

//将字符串首字母转大写
export const titleUpperCase = function (name) {
  if (!name) {
    return ''
  }
  return name.replace(/^./, function (match) {
    return match.toUpperCase()
  })
}
