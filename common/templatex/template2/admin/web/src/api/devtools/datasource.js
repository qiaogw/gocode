import { api } from 'src/boot/axios'

// gensource 列表
export function listDatasource(data) {
  return api({
    url: '/gencode/gensource/list',
    method: 'post',
    data: data,
  })
}

// gensource create
export function createDatasource(data) {
  return api({
    url: '/gencode/gensource/create',
    method: 'post',
    data: data,
  })
}
// gensource delete
export function deleteDatasource(data) {
  return api({
    url: '/gencode/gensource/delete',
    method: 'post',
    data: data,
  })
}
// gensource update
export function updateDatasource(data) {
  return api({
    url: '/gencode/gensource/update',
    method: 'post',
    data: data,
  })
}
// gensource test
export function testDatasource(data) {
  return api({
    url: '/gencode/gensource/test',
    method: 'post',
    data: data,
  })
}
// gensource getTables
export function getTables(data) {
  return api({
    url: '/gencode/gensource/getTables',
    method: 'post',
    data: data,
  })
}

// gensource getTables
export function getTable(data) {
  return api({
    url: '/gencode/gensource/getTable',
    method: 'post',
    data: data,
  })
}

// gensource getColumns
export function getColumns(data) {
  return api({
    url: '/gencode/gensource/getColumns',
    method: 'post',
    data: data,
  })
}

// gensource previewCode
export function previewCode(data) {
  return api({
    url: '/gencode/gensource/previewCode',
    method: 'post',
    data: data,
  })
}

// gensource genCode
export function genCode(data) {
  return api({
    url: '/gencode/gensource/genGetCode',
    method: 'post',
    data: data,
  })
}

// gensource export
export function excelExport(data) {
  return api({
    url: '/gencode/gensource/export',
    method: 'post',
    data: data,
    responseType: 'arraybuffer',
  })
}

// gensource import
export function excelImport(data) {
  return api({
    url: '/gencode/gensource/import',
    method: 'post',
    data: data,
  })
}

// gensource 生成目录数据库
export function gendb(data) {
  return api({
    url: '/gencode/gensource/gendb',
    method: 'post',
    data: data,
  })
}

// genCoverCode 覆盖本机代码
export function genCoverCode(data) {
  return api({
    url: '/gencode/gensource/genCoverCode',
    method: 'post',
    data: data,
  })
}
