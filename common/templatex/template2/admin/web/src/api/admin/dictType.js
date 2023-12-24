import { api } from "src/boot/axios";

// dicttype 列表
export function listDict(data) {
  return api({
    url: "admin/dicttype/list",
    method: "post",
    data: data,
  });
}

// dicttype create
export function createDict(data) {
  return api({
    url: "admin/dicttype/create",
    method: "post",
    data: data,
  });
}

// dicttype update
export function updateDict(data) {
  return api({
    url: "admin/dicttype/update",
    method: "post",
    data: data,
  });
}

// dicttype delete
export function deleteDict(data) {
  return api({
    url: "admin/dicttype/delete",
    method: "post",
    data: data,
  });
}
