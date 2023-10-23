import { api } from "src/boot/axios";

// {{.TableComment}} 列表
export function list{{.Table}}(data) {
  return api({
    url: "{{.TableUrl}}/list",
    method: "post",
    data: data,
  });
}

// {{.TableComment}} 获取
export function get{{.Table}}(data) {
  return api({
    url: "{{.TableUrl}}/get",
    method: "post",
    data: data,
  });
}

// {{.TableComment}} 创建
export function create{{.Table}}(data) {
  return api({
    url: "{{.TableUrl}}/create",
    method: "post",
    data: data,
  });
}
// {{.TableComment}} 删除
export function delete{{.Table}}(data) {
  return api({
    url: "{{.TableUrl}}/delete",
    method: "post",
    data: data,
  });
}
// {{.TableComment}} 更新
export function update{{.Table}}(data) {
  return api({
    url: "{{.TableUrl}}/update",
    method: "post",
    data: data,
  });
}