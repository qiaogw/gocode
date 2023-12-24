import { api } from 'src/boot/axios'

// 用户 列表
export function listUser(data) {
  return api({
    url: 'admin/user/list',
    method: 'post',
    data: data,
  })
}

// 用户 获取
export function getUser(data) {
  return api({
    url: 'admin/user/get',
    method: 'post',
    data: data,
  })
}

// 用户 创建
export function createUser(data) {
  return api({
    url: 'admin/user/create',
    method: 'post',
    data: data,
  })
}
// 用户 删除
export function deleteUser(data) {
  return api({
    url: 'admin/user/delete',
    method: 'post',
    data: data,
  })
}
// 用户 更新
export function updateUser(data) {
  return api({
    url: 'admin/user/update',
    method: 'post',
    data: data,
  })
}

// 用户 切换角色
export function setMeRole(data) {
  return api({
    url: 'admin/user/setMeRole',
    method: 'post',
    data: data,
  })
}

// 用户 重置密码
export function resetPassword(data) {
  return api({
    url: 'admin/user/resetPassword',
    method: 'post',
    data: data,
  })
}

// 用户 修改密码
export function setPassword(data) {
  return api({
    url: 'admin/user/setPassword',
    method: 'post',
    data: data,
  })
}

// 列表 用户 无此角色
export function listNoUser(data) {
  return api({
    url: 'admin/user/listNoUser',
    method: 'post',
    data: data,
  })
}

// 列表 用户树
export function treeUser(data) {
  return api({
    url: 'admin/user/tree',
    method: 'post',
    data: data,
  })
}
