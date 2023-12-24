import { SessionStorage } from 'quasar'
import { useDictionaryStore } from 'src/stores/dict'
import XEUtils from 'xe-utils'
export async function DictOptions() {
  const dictStore = useDictionaryStore()
  const dicts = await dictStore.GetDicts()
  return dicts
}

// 回显数据字典
export function getDictLabel(datas, name) {
  let dict = XEUtils.find(datas, (item) => item.value + '' === name + '')
  if (!dict) {
    return name
  }
  return dict.label
}
// 获取字典
export async function getDict(name) {
  const dictionaryStore = useDictionaryStore()
  await dictionaryStore.GetSubDict(name)
  return dictionaryStore.dictionaryMap[name]
}
// 获取列表字典
export function getDictList(datas, name) {
  let dict = XEUtils.map(datas, (item) => item[name])
  if (!dict) {
    return
  }
  return dict
}

export function getOptionsByList(datas, label, value) {
  const data = []
  let k = label || 'id'
  let v = value || 'name'
  if (datas && datas.length > 0) {
    datas.forEach((e) => {
      data.push({
        label: e[k],
        value: e[v],
      })
    })
    return data
  }
}
