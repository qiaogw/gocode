import { api } from 'src/boot/axios'

// 职务 列表
export function listPost(data) {
  return api({
    url: 'admin/post/list',
    method: 'post',
    data: data,
  })
}

// 职务 获取
export function getPost(data) {
  return api({
    url: 'admin/post/get',
    method: 'post',
    data: data,
  })
}

// 职务 创建
export function createPost(data) {
  return api({
    url: 'admin/post/create',
    method: 'post',
    data: data,
  })
}
// 职务 删除
export function deletePost(data) {
  return api({
    url: 'admin/post/delete',
    method: 'post',
    data: data,
  })
}
// 职务 更新
export function updatePost(data) {
  return api({
    url: 'admin/post/update',
    method: 'post',
    data: data,
  })
}
