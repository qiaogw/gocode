import { api } from 'src/boot/axios'

//  列表
export function listLog(data) {
  return api({
    url: 'task/log/list',
    method: 'post',
    data: data,
  })
}

//  获取
export function getLog(data) {
  return api({
    url: 'task/log/get',
    method: 'post',
    data: data,
  })
}

//  创建
export function createLog(data) {
  return api({
    url: 'task/log/create',
    method: 'post',
    data: data,
  })
}
//  删除
export function deleteLog(data) {
  return api({
    url: 'task/log/delete',
    method: 'post',
    data: data,
  })
}

//  删除
export function deleteLogList(data) {
  return api({
    url: 'task/log/deleteList',
    method: 'post',
    data: data,
  })
}

//  更新
export function updateLog(data) {
  return api({
    url: 'task/log/update',
    method: 'post',
    data: data,
  })
}

//  更新
export function stopTask(data) {
  return api({
    url: 'task/log/stop',
    method: 'post',
    data: data,
  })
}
