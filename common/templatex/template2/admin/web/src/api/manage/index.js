import { api } from "src/boot/axios";
import { Notify } from "quasar";

export function getAction(url, params) {
  return api({
    url: url,
    method: "get",
    params: params,
  });
}

export function postAction(url, params) {
  return api({
    url: url,
    method: "post",
    data: params,
  });
}

export function downFile(url, params) {
  return api({
    url: url,
    method: "post",
    data: params,
    responseType: "arraybuffer",
  });
}

export async function downloadAction(url, fileName, params) {
  const res = await downFile(url, params);
  const data = res.data;
  if (!data || data.size === 0) {
    Notify.create({
      type: "negative",
      message: "文件下载失败！",
    });
    return;
  }
  if (typeof window.navigator.msSaveBlob !== "undefined") {
    window.navigator.msSaveBlob(new Blob([data]), fileName);
  } else {
    let urlHref = window.URL.createObjectURL(new Blob([data]));
    let link = document.createElement("a");
    link.style.display = "none";
    link.href = urlHref;
    link.setAttribute("download", fileName);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link); //下载完成移除元素
    window.URL.revokeObjectURL(urlHref); //释放掉blob对象
  }
}
