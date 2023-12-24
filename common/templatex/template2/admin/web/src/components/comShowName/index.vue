<template>
  <div>
    <span>
      {{ trueName }}
    </span>
  </div>
</template>

<script setup>
import { useQuasar } from 'quasar'

import { computed, toRefs } from 'vue'
import useCommon from 'src/composables/useCommon'
import { useUserStore } from 'src/stores/user'
const userStore = useUserStore()

const { subDefaultUsername } = useCommon()
const $q = useQuasar()
const props = defineProps({
  customNameString: {
    type: String,
    required: false,
    default: '',
  },
  customNameObject: {
    type: Object,
    required: false,
    default: () => {
      return {}
    },
  },
  showMyName: {
    type: Boolean,
    required: false,
    default: false,
  },
})
const { customNameString, customNameObject, showMyName } = toRefs(props)

const trueName = computed(() => {
  if (customNameString.value !== '') {
    // 自定义名字
    return customNameString.value
  } else if (JSON.stringify(customNameObject.value) !== '{}') {
    // 其他用户的名字
    if (customNameObject.value.nickname) {
      return customNameObject.value.nickname
    } else if (customNameObject.value.real_name) {
      return customNameObject.value.real_name
    } else {
      return customNameObject.value.username
    }
  } else if (showMyName.value) {
    const user = userStore.GetInfo()
    if (!user) {
      return
    }
    const nickname = user.nickName
    const realName = user.realName
    if (nickname) {
      return nickname
    } else {
      return realName
    }
  } else {
    return subDefaultUsername
  }
})
</script>
