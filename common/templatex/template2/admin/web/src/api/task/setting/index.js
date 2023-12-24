import { api } from 'src/boot/axios'

//  列表
export function listSetting(data) {
  return api({
    url: 'task/setting/list',
    method: 'post',
    data: data,
  })
}

//  获取
export function getSetting(data) {
  return api({
    url: 'task/setting/get',
    method: 'post',
    data: data,
  })
}

//  创建
export function createSetting(data) {
  return api({
    url: 'task/setting/create',
    method: 'post',
    data: data,
  })
}
//  删除
export function deleteSetting(data) {
  return api({
    url: 'task/setting/delete',
    method: 'post',
    data: data,
  })
}
//  更新
export function updateSetting(data) {
  return api({
    url: 'task/setting/update',
    method: 'post',
    data: data,
  })
}

// mail
export function getTaskMailServer(data) {
  return api({
    url: 'task/setting/getmail',
    method: 'post',
    data: data,
  })
}
// updateMail
export function updateMail(data) {
  return api({
    url: 'task/setting/updatemail',
    method: 'post',
    data: data,
  })
}

// createMailUser
export function createMailUser(data) {
  return api({
    url: 'task/setting/createmailuser',
    method: 'post',
    data: data,
  })
}
// deleteMailUser
export function deleteMailUser(data) {
  return api({
    url: 'task/setting/deletemailuser',
    method: 'post',
    data: data,
  })
}
