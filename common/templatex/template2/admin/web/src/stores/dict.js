import { defineStore } from 'pinia'
import { listDict } from 'src/api/admin/dictType'
import { ArrayToTree } from 'src/utils/arrayAndTree'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const SetSubDict = async () => {
    const queryReq = {
      pageIndex: 0,
      pageSize: 9999,
    }
    const res = await listDict(queryReq)
    if (res) {
      let dictDetail = res.list
      let dict = {}
      for (let d of dictDetail) {
        dict[d.type] = d.dictDatas
      }
      dictionaryMap.value = dict
    }
  }
  const GetSubDict = async (type) => {
    if (
      dictionaryMap.value &&
      dictionaryMap.value[type] &&
      dictionaryMap.value[type].length
    ) {
      return dictionaryMap.value[type]
    } else {
      await SetSubDict()
      return dictionaryMap.value[type]
    }
  }
  const GetDicts = async () => {
    if (dictionaryMap.value) {
      return dictionaryMap.value
    } else {
      await SetSubDict()
      return dictionaryMap.value
    }
  }
  const clear = async () => {
    dictionaryMap.value = undefined
  }
  return {
    dictionaryMap,
    GetSubDict,
    SetSubDict,
    clear,
    GetDicts,
  }
})
