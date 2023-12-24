import { api } from "src/boot/axios";

// dictdata 列表
export function listDict(data) {
  return api({
    url: "admin/dictdata/list",
    method: "post",
    data: data,
  });
}

// dictdata create
export function createDict(data) {
  return api({
    url: "admin/dictdata/create",
    method: "post",
    data: data,
  });
}

// dictdata update
export function updateDict(data) {
  return api({
    url: "admin/dictdata/update",
    method: "post",
    data: data,
  });
}

// dictdata delete
export function deleteDict(data) {
  return api({
    url: "admin/dictdata/delete",
    method: "post",
    data: data,
  });
}
