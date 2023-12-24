<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <q-layout style="overflow-x: hidden">
    <q-page-container>
      <q-form @submit="submitForm" @reset="onReset" class="q-gutter-md">
        <q-input
          v-model="myForm.username"
          outlined
          dense
          no-error-icon
          placeholder="请输入用户名"
          :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]"
        >
        </q-input>
        <q-input
          v-model="myForm.password"
          filled
          :type="isPwd ? 'password' : 'text'"
          placeholder="请输入密码"
        >
          <template v-slot:append>
            <q-icon
              :name="isPwd ? 'visibility_off' : 'visibility'"
              class="cursor-pointer"
              @click="isPwd = !isPwd"
            />
          </template>
        </q-input>
        <div class="vPicBox">
          <q-input
            v-model="myForm.captcha"
            placeholder="请输入验证码"
            style="width: 60%"
          />
          <div class="vPic">
            <img
              v-if="picPath"
              :src="picPath"
              alt="请输入验证码"
              @click="getCaptcha()"
            />
          </div>
        </div>
        <div>
          <q-btn
            label="登 录"
            type="submit"
            size="large"
            style="width: 46%; margin-left: 8%"
          ></q-btn>
          <q-btn
            label="Reset"
            type="reset"
            color="primary"
            flat
            class="q-ml-sm"
          />
        </div>
      </q-form>
      <!-- <q-btn type="submit" style="width: 46%" size="large" @click="checkInit"
        >前往初始化</q-btn
      > -->
    </q-page-container>
  </q-layout>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'login',
}
</script>

<script setup>
import { captcha } from 'src/api/admin/auth'
// import { checkDB } from "@/api/initdb";
// import BottomInfo from "@/view/layout/bottomInfo/bottomInfo.vue";
import { onMounted, reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from 'src/stores/user'

const router = useRouter()
const route = useRoute()
const isPwd = ref(true)
const picPath = ref('')
const myForm = reactive({
  username: 'admin',
  password: '123456',
  mobile: '',
  captcha: '',
  captchaId: '',
})

// 获取验证码
const getCaptcha = () => {
  captcha().then((res) => {
    if (res) {
      picPath.value = res.picPath
      myForm.captchaId = res.captchaId
    }
  })
}
onMounted(() => {
  getCaptcha()
})

const rules = reactive({
  username: [{ required: true, trigger: 'blur' }],
  password: [{ required: true, trigger: 'blur' }],
  captcha: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
// const login = async () => {
//   return await userStore.Login(myForm);
// };
const submitForm = async () => {
  myForm.mobile = myForm.username
  const res = await userStore.Login(myForm)
  getCaptcha()
  if (res) {
    router.push(route.query.redirect || '/')
    // delete myForm[captcha];
    myForm.captcha = ''
  } else {
    myForm.captcha = ''
  }
}

// 跳转初始化
const checkInit = async () => {
  //   const res = await checkDB();
  //   if (res.code === 0) {
  //     if (res.data?.needInit) {
  //       userStore.NeedInit();
  //       router.push({ name: "Init" });
  //     } else {
  //       ElMessage({
  //         type: "info",
  //         message: "已配置数据库信息，无法初始化",
  //       });
  //     }
  //   }
}
const onReset = () => {
  Object.keys(myForm).map((key) => {
    delete myForm[key]
  })
  getCaptcha()
}
</script>
