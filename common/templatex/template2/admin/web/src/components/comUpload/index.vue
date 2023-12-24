<template>
  <q-btn icon="upload" :rounded="rounded" flat no-wrap glossy color="primary">
    <q-tooltip>{{ title }}</q-tooltip>
    <q-popup-proxy ref="uplaodPip">
      <q-uploader
        color="teal"
        flat
        bordered
        style="max-width: 300px"
        :multiple="multiple"
        :max-file-size="maxFileSize * 1024 * 1024"
        :factory="factoryFn"
        :accept="fileType"
        @uploaded="uploaded"
        @failed="failed"
        @finish="finish"
        @rejected="rejected"
      />
    </q-popup-proxy>
  </q-btn>
</template>

<script setup>
import { ref, computed, onBeforeUnmount, getCurrentInstance, toRefs } from 'vue'
import { useUserStore } from 'src/stores/user'
import { useQuasar, QSpinnerFacebook } from 'quasar'
const $q = useQuasar()
let { proxy } = getCurrentInstance()
const uplaodPip = ref(null)

const props = defineProps({
  title: String,
  rounded: Boolean,
  multiple: Boolean,
  uploadUrl: String,
  maxFileSize: {
    type: Number,
    required: false,
    default: 10,
  },
  fileType: {
    type: String,
    required: false,
    default: '*',
  },
})
const { title, multiple, rounded, uploadUrl, maxFileSize, fileType } =
  toRefs(props)
const userStore = useUserStore()
const token = computed(() => userStore.GetToken())

const emit = defineEmits(['upload'])

const factoryFn = (files) => {
  $q.loading.show({
    spinner: QSpinnerFacebook,
    spinnerColor: 'green',
    spinnerSize: 140,
    backgroundColor: 'blue',
    message: '文件正在上传中,请稍候...',
    messageColor: 'yellow',
  })
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
  $q.loading.hide()
}
const failed = (info) => {
  const res = JSON.parse(info.xhr.response)
  $q.loading.hide()
  proxy.$error(props.title + res.msg)
}

const rejected = (rejectedEntries) => {
  $q.loading.hide()
  proxy.$warn(props.title + '文件重复或大小/类型不被允许')
}

const finish = () => {
  $q.loading.hide()
  uplaodPip.value.hide()
  emit('upload')
}

const checkFileSize = (files) => {
  return files.filter((file) => file.size < 2048)
}

const checkFileType = (files) => {
  return files.filter((file) => file.type === 'image/png')
}
</script>
