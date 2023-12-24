import { api } from "src/boot/axios";

// genpkg 列表
export function listPkg(data) {
  return api({
    url: "/gencode/genpkg/list",
    method: "post",
    data: data,
  });
}

// genpkg create
export function createPkg(data) {
  return api({
    url: "/gencode/genpkg/create",
    method: "post",
    data: data,
  });
}
// genpkg delete
export function deletePkg(data) {
  return api({
    url: "/gencode/genpkg/delete",
    method: "post",
    data: data,
  });
}
// genpkg update
export function updatePkg(data) {
  return api({
    url: "/gencode/genpkg/update",
    method: "post",
    data: data,
  });
}

// genPkgCode 生成代码
export function genPkgCode(data) {
  return api({
    url: "/gencode/genpkg/gen",
    method: "post",
    data: data,
  });
}
