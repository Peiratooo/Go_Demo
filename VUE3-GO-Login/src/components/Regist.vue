<template>
    <div class="loginCard">
        <div class="input">
            <div class="inputTitle">用户名:</div>
            <input type="text" v-model="user.name">
        </div>
        <div class="input">
            <div class="inputTitle">邮箱:</div>
            <input type="text" v-model="user.email">
        </div>

        <div class="input">
            <div class="inputTitle">密码:</div>
            <input type="password" v-model="user.password">
        </div>
        <div class="button" @click="regist">注 册</div>

    </div>
    <div class="subTitle">已有账号？<span @click="router.push('/login')">登 录</span></div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { reactive } from 'vue'
import axios from 'axios'

const request = axios.create({
    baseURL: '/api/entrance/register',
    timeout: 1000
})
const router = useRouter()
const user = reactive({
    name: "",
    email: "",
    password: ""
})
function regist() {
    // console.log(user);
    if (!user.email || !user.name || !user.password) {
        alert("请完善用户信息！")
    } else {
        request.post('', user).then(res => {
            console.log(res);
            const { status } = res
            if (status == 200) {
                if (res.data.msg == "OK") {
                    window.localStorage.removeItem('token');
                    router.push({
                        name: "login",
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