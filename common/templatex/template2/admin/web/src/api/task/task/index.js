import { api } from 'src/boot/axios'

//  列表
export function listTask(data) {
  return api({
    url: 'task/task/list',
    method: 'post',
    data: data,
  })
}

//  获取
export function getTask(data) {
  return api({
    url: 'task/task/get',
    method: 'post',
    data: data,
  })
}

//  创建
export function createTask(data) {
  return api({
    url: 'task/task/create',
    method: 'post',
    data: data,
  })
}
//  删除
export function deleteTask(data) {
  return api({
    url: 'task/task/delete',
    method: 'post',
    data: data,
  })
}
//  更新
export function updateTask(data) {
  return api({
    url: 'task/task/update',
    method: 'post',
    data: data,
  })
}

//  手动运行
export function run(data) {
  return api({
    url: 'task/task/run',
    method: 'post',
    data: data,
  })
}

//  获取子任务
export function getChild(data) {
  return api({
    url: 'task/task/getChild',
    method: 'post',
    data: data,
  })
}
