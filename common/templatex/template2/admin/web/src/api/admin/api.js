import { api } from "src/boot/axios";

// api 列表
export function listApi(data) {
  return api({
    url: "/admin/api/list",
    method: "post",
    data: data,
  });
}

// api create
export function createApi(data) {
  return api({
    url: "/admin/api/create",
    method: "post",
    data: data,
  });
}
// api delete
export function deleteApi(data) {
  return api({
    url: "/admin/api/delete",
    method: "post",
    data: data,
  });
}
// api update
export function updateApi(data) {
  return api({
    url: "/admin/api/update",
    method: "post",
    data: data,
  });
}
// 提取-用户api-根据角色
export function getApiByRole(data) {
  return api({
    url: "/admin/api/getApiByRole",
    method: "post",
    data: data,
  });
}
