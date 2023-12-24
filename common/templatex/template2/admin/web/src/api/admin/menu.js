import { api } from "src/boot/axios";

// menu 列表
export function listMenu(data) {
  return api({
    url: "admin//menu/list",
    method: "post",
    data: data,
  });
}

// menu create
export function createMenu(data) {
  return api({
    url: "admin//menu/create",
    method: "post",
    data: data,
  });
}
// menu delete
export function deleteMenu(data) {
  return api({
    url: "admin//menu/delete",
    method: "post",
    data: data,
  });
}
// menu update
export function updateMenu(data) {
  return api({
    url: "admin//menu/update",
    method: "post",
    data: data,
  });
}
// 提取-用户菜单-根据角色
export function getMenuByRole(data) {
  return api({
    url: "admin//menu/getMenuByRole",
    method: "post",
    data: data,
  });
}
