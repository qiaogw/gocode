import { api } from 'src/boot/axios'

// 主机节点 列表
export function listNode(data) {
  return api({
    url: 'task/node/list',
    method: 'post',
    data: data,
  })
}

// 主机节点 获取
export function getNode(data) {
  return api({
    url: 'task/node/get',
    method: 'post',
    data: data,
  })
}

// 主机节点 创建
export function createNode(data) {
  return api({
    url: 'task/node/create',
    method: 'post',
    data: data,
  })
}
// 主机节点 删除
export function deleteNode(data) {
  return api({
    url: 'task/node/delete',
    method: 'post',
    data: data,
  })
}
// 主机节点 更新
export function updateNode(data) {
  return api({
    url: 'task/node/update',
    method: 'post',
    data: data,
  })
}

// 主机节点 连接测试
export function pingNode(data) {
  return api({
    url: 'task/node/ping',
    method: 'post',
    data: data,
  })
}
