import { api } from "src/boot/axios";

// 查询用户列表
export function captcha(data) {
  return api({
    url: "/auth/captcha",
    method: "post",
    data: data,
  });
}
// 查询用户列表
export function login(data) {
  return api({
    url: "/auth/login",
    method: "post",
    data: data,
  });
}

// 查询用户自己
export function getMe(data) {
  return api({
    url: "/admin/user/getMe",
    method: "post",
    data: data,
  });
}

// 查询用户自己菜单
export function getMenu(data) {
  return api({
    url: "/admin/menu/getMenu",
    method: "post",
    data: data,
  });
}
