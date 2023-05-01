<template>
    <div class="loginCard">
        <div class="input">
            <div class="inputTitle">账号/邮箱:</div>
            <input type="text" placeholder="请输入账号/邮箱" v-model="user.uidoremail">
        </div>
        <div class="input">
            <div class="inputTitle">密码:</div>
            <input type="password" placeholder="请输入密码" v-model="user.password">
        </div>
        <div class="button" @click="login">登 录</div>
    </div>
    <div class="subTitle">没有账号？<span @click="router.push('/regist')">注 册</span></div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { reactive } from 'vue'
import axios from 'axios'
const router = useRouter()
const user = reactive({
    uidoremail: "",
    password: ""
})

const request = axios.create({
    baseURL: '/api/entrance/login',
    timeout: 1000
})

function login() {
    if (!user.uidoremail || !user.password) {
        alert("请完善用户信息！")
    } else {
        request.post('', user).then(res => {
            console.log(res);
            const { status } = res
            if (status == 200) {
                if (res.data.msg == "OK") {
                    window.localStorage.setItem('Token', res.data.token);
                    router.push({
                        path: "/"
                    })
                } else {
                    alert(res.data.msg)
                }

            } else {
                alert(res.data.msg)
            }
        })
    }
}
</script>

<style></style>