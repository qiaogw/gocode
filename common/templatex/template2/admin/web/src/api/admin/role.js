import { api } from 'src/boot/axios'

// role 列表
export function listRole(data) {
  return api({
    url: 'admin/role/list',
    method: 'post',
    data: data,
  })
}

// role get
export function getRole(data) {
  return api({
    url: 'admin/role/get',
    method: 'post',
    data: data,
  })
}

// role create
export function createRole(data) {
  return api({
    url: 'admin/role/create',
    method: 'post',
    data: data,
  })
}
// role delete
export function deleteRole(data) {
  return api({
    url: 'admin/role/delete',
    method: 'post',
    data: data,
  })
}
// role update
export function updateRole(data) {
  return api({
    url: 'admin/role/update',
    method: 'post',
    data: data,
  })
}

// RoleMenu update
export function setRoleMenu(data) {
  return api({
    url: 'admin/role/setRoleMenu',
    method: 'post',
    data: data,
  })
}

// RoleApi update
export function setRoleApi(data) {
  return api({
    url: 'admin/role/setRoleApi',
    method: 'post',
    data: data,
  })
}

// RoleApi 角色添加用户
export function setRoleUser(data) {
  return api({
    url: 'admin/role/setRoleUsers',
    method: 'post',
    data: data,
  })
}
