import { api } from 'src/boot/axios'

// gentable 列表
export function listGenTable(data) {
  return api({
    url: '/gencode/gentable/list',
    method: 'post',
    data: data,
  })
}
// 表模块 获取
export function getGenTable(data) {
  return api({
    url: 'gencode/gentable/get',
    method: 'post',
    data: data,
  })
}

// gentable create
export function createGenTable(data) {
  return api({
    url: '/gencode/gentable/create',
    method: 'post',
    data: data,
  })
}
// gentable delete
export function deleteGenTable(data) {
  return api({
    url: '/gencode/gentable/delete',
    method: 'post',
    data: data,
  })
}
// gentable update
export function updateGenTable(data) {
  return api({
    url: '/gencode/gentable/update',
    method: 'post',
    data: data,
  })
}
