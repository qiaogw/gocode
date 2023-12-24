<template>
  <q-btn icon="downloading" :rounded="rounded" no-wrap glossy color="primary">
    <q-tooltip>{{ title }}</q-tooltip>
  </q-btn>
</template>

<script setup>
import { ref, computed, onBeforeUnmount, getCurrentInstance, toRefs } from 'vue'
import { useUserStore } from 'src/stores/user'
import { useQuasar } from 'quasar'
const $q = useQuasar()
let { proxy } = getCurrentInstance()

const props = defineProps({
  title: String,
  multiple: Boolean,
  downloadUrl: String,
  rounded: Boolean,
})
const { title, multiple, rounded, uploadUrl, maxFileSize, fileType } =
  toRefs(props)
const userStore = useUserStore()
const token = computed(() => userStore.GetToken())

const emit = defineEmits(['upload'])

const factoryFn = (files) => {
  return new Promise((resolve) => {
    // simulating a delay of 2 seconds
    setTimeout(() => {
      resolve({
        url: uploadUrl.value,
        headers: [{ name: 'Authorization', value: token.value }],
        fieldName: 'upFile',
        method: 'POST',
      })
    }, 2000)
  })
}

const uploaded = (info) => {
  const res = JSON.parse(info.xhr.response)
  if (res.code === 0) {
    proxy.$info(props.title + '成功！')
  } else {
    proxy.$warn(props.title + res.msg)
  }
}
const failed = (info) => {
  const res = JSON.parse(info.xhr.response)
  proxy.$error(props.title + res.msg)
}

const rejected = (rejectedEntries) => {
  proxy.$warn(props.title + '文件重复或大小/类型不被允许')
}

const finish = () => {
  emit('upload')
}

const checkFileSize = (files) => {
  return files.filter((file) => file.size < 2048)
}

const checkFileType = (files) => {
  return files.filter((file) => file.type === 'image/png')
}
</script>
