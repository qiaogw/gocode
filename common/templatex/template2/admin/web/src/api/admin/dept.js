import { api } from 'src/boot/axios'

// 部门 列表
export function listDept(data) {
  return api({
    url: 'admin/dept/list',
    method: 'post',
    data: data,
  })
}

// 部门 获取
export function getDept(data) {
  return api({
    url: 'admin/dept/get',
    method: 'post',
    data: data,
  })
}

// 部门 创建
export function createDept(data) {
  return api({
    url: 'admin/dept/create',
    method: 'post',
    data: data,
  })
}
// 部门 删除
export function deleteDept(data) {
  return api({
    url: 'admin/dept/delete',
    method: 'post',
    data: data,
  })
}
// 部门 更新
export function updateDept(data) {
  return api({
    url: 'admin/dept/update',
    method: 'post',
    data: data,
  })
}
