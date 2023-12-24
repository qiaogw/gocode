<template>
  <q-layout style="overflow-x: hidden">
    <q-page-container>
      <q-card class="my-card fixed-center" flat bordered>
        <q-card-section>
          <div class="text-h6">Our Changing Planet</div>
          <div class="text-subtitle2">by John Doe</div>
        </q-card-section>

        <q-tabs v-model="tab" class="text-teal">
          <q-tab label="账号登录" name="one" />
          <q-tab label="微信登录" name="two" />
        </q-tabs>

        <q-separator />
        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="one">
            <q-form
              ref="loginForm"
              @submit="submitForm"
              @reset="onReset"
              class="q-gutter-md login-form"
              label-position="left"
            >
              <q-input
                v-model="myForm.username"
                outlined
                dense
                no-error-icon
                placeholder="请输入用户名"
                :rules="[(val) => (val && val.length > 0) || $t('NeedInput')]"
              />
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
          </q-tab-panel>

          <q-tab-panel name="two">
            <div id="login_wechat"></div>
            <q-btn
              label="微信"
              type="submit"
              size="large"
              @click="getWechat"
              style="width: 46%; margin-left: 8%"
            ></q-btn>
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import SocialSign from './sign.vue'
import { getSns } from 'src/api/usercenter/usersns'
import { captcha } from 'src/api/admin/auth'
// import 'https://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js'
// import { checkDB } from "@/api/initdb";
// import BottomInfo from "@/view/layout/bottomInfo/bottomInfo.vue";
import { onMounted, reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from 'src/stores/user'

const router = useRouter()
const route = useRoute()
const isPwd = ref(true)
const picPath = ref('')
const showAnimate = ref(false)
const tab = ref('one')

const wechat = reactive({
  showQR: true,
  loginType: 'QRcode',
  showReg: false,
  showLogin: true,
  appid: '',
  scope: 'snsapi_login',
  redirect_uri: '',
})

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
// 获取微信信息
const getWechat = async () => {
  let res = await getSns()
  wechat.appid = res.appId
  wechat.redirect_uri = res.loginUrl
  // window.location.href = res.loginUrl
  // console.log('wechat', wechat)
  const s = document.createElement('script')
  s.type = 'text/javascript'
  s.src = 'https://res.wx.qq.com/connect/zh_CN/htmledition/js/wxLogin.js'
  const wxElement = document.body.appendChild(s)
  // const uri = `${window.location.origin}callback/wx/` // 这里是你的回调uri
  wxElement.onload = () => {
    const obj = new WxLogin({
      self_redirect: false,
      id: 'login_wechat', // 需要显示的容器id
      appid: wechat.appid, // appid wx*******
      scope: 'snsapi_base', // 网页默认即可
      redirect_uri: wechat.loginUrl, // 授权成功后回调的url
      state: Math.ceil(Math.random() * 1000), // 可设置为简单的随机数加session用来校验
      style: 'white', // 提供"black"、"white"可选。二维码的样式
      href: ``, // 外部css文件url，需要https
    })
    if (!obj) {
      console.error('wx-error')
    }
    //
  }
}

onMounted(async () => {
  getCaptcha()
  // await getWechat()
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

const reset = () => {
  showAnimate.value = false
}
</script>
<style lang="sass" scoped>
.my-card
  width: 100%
  max-width: 550px
</style>
