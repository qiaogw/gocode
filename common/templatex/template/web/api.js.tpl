import { api } from "src/boot/axios"

// {{.TableComment}} 列表
export function list{{.Table}}(data) {
  return api({
    url: "{{.Db}}/{{.TableUrl}}/list",
    method: "post",
    data: data,
  })
}

// {{.TableComment}} 获取
export function get{{.Table}}(data) {
  return api({
    url: "{{.Db}}/{{.TableUrl}}/get",
    method: "post",
    data: data,
  })
}

// {{.TableComment}} 创建
export function create{{.Table}}(data) {
  return api({
    url: "{{.Db}}/{{.TableUrl}}/create",
    method: "post",
    data: data,
  })
}
// {{.TableComment}} 删除
export function delete{{.Table}}(data) {
  return api({
    url: "{{.Db}}/{{.TableUrl}}/delete",
    method: "post",
    data: data,
  })
}
// {{.TableComment}} 批量删除
export function deleteList{{.Table}}(data) {
  return api({
  url: "{{.Db}}/{{.TableUrl}}/deleteList",
  method: "post",
  data: data,
})
}
// {{.TableComment}} 更新
export function update{{.Table}}(data) {
  return api({
    url: "{{.Db}}/{{.TableUrl}}/update",
    method: "post",
    data: data,
  })
}

{{- if .IsFlow }}
// {{.TableComment}} 工作流审批转换
export function trigger(data) {
  return api({
    url: '{{.Db}}/{{.TableUrl}}/trigger',
    method: 'post',
    data: data,
  })
}
{{- end}}
